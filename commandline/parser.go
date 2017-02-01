package commandline

import (
	"fmt"
	"github.com/maxmeyer/letter-generator-go/letter_generator"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"strings"
)

type Cli struct{}

func (p *Cli) Run(args []string, config letter_generator.Config, current_directory string) error {

	appMetadata := letter_generator.AppMetadata{
		Version: "0.0.1",
		License: "MIT",
		Authors: []letter_generator.AppAuthor{
			letter_generator.AppAuthor{
				Name:  "Dennis Günnewig",
				Email: "dev@fedux.org",
			},
		},
	}

	app := cli.NewApp()
	app.Name = "letter-generator"
	app.Version = appMetadata.Version

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose, V",
			Usage: "activate verbose logging",
		},
		cli.BoolFlag{
			Name:  "show-config, C",
			Usage: "Show configuration",
		},
	}

	app.Action = func(c *cli.Context) error {
		parse_global_options(c, config)

		err := build(config)

		if err != nil {
			return err
		}

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "initialize current directory",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "verbose, V",
					Usage: "activate verbose logging",
				},
			},
			Action: func(c *cli.Context) error {
				parse_global_options(c, config)

				err := initialize(current_directory, config)

				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build letters based on information in current directory",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "verbose, V",
					Usage: "activate verbose logging",
				},
			},
			Action: func(c *cli.Context) error {
				parse_global_options(c, config)

				err := build(config)

				if err != nil {
					return err
				}

				return nil
			},
		},
	}

	app.Run(os.Args)

	return nil
}

func build(config letter_generator.Config) error {
	builder := LetterBuilder{}
	err := builder.Build(config)

	if err != nil {
		return err
	}

	return nil
}

func initialize(dir string, config letter_generator.Config) error {
	initializer := Initializer{}
	err := initializer.Init(dir, config)

	if err != nil {
		return err
	}

	return nil
}

func parse_global_options(c *cli.Context, config letter_generator.Config) {
	if c.Bool("verbose") == true {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.WithFields(log.Fields{
		"verbose":     c.Bool("verbose"),
		"show-config": c.Bool("show-config"),
	}).Info("Parsing commandline options")

	if c.Bool("show-config") == true {
		fmt.Print(strings.Join(config.ToString(), "\n"))
		fmt.Print("\n")
		os.Exit(0)
	}
}
