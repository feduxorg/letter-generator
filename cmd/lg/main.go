package main

import (
	"github.com/maxmeyer/letter-generator-go/commandline"
	"github.com/maxmeyer/letter-generator-go/letter_generator"
	lgos "github.com/maxmeyer/letter-generator-go/os"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr, could also be a file.
	// log.SetOutput(os.Stdout)
}

func main() {
	current_directory, err := os.Getwd()

	if err != nil {
		log.WithFields(log.Fields{
			"msg":    err.Error(),
			"status": "failure",
		}).Fatal("Getting current directory")

		os.Exit(1)
	}

	log.WithFields(log.Fields{
		"path":   current_directory,
		"status": "success",
	}).Debug("Getting current directory")

	home_directory, err := lgos.HomeDirectory()

	if err != nil {
		log.WithFields(log.Fields{
			"msg":    err.Error(),
			"status": "failure",
		}).Fatal("Getting home directory of current user")

		os.Exit(1)
	}

	log.WithFields(log.Fields{
		"path":   home_directory,
		"status": "success",
	}).Debug("Getting home directory of current user")

	config := letter_generator.Config{}
	config.RemoteSources = []string{filepath.Join(home_directory, ".local/share/letter-template/config.git"), "git@gitlab.com:maxmeyer/letter-template.git"}
	config.ConfigDirectory = ".lg"
	config.RecipientsFile = filepath.Join(current_directory, config.ConfigDirectory, "data/to.json")
	config.MetadataFile = filepath.Join(current_directory, config.ConfigDirectory, "data/metadata.json")
	config.SenderFile = filepath.Join(current_directory, config.ConfigDirectory, "data/from.json")
	config.TemplateFile = filepath.Join(current_directory, config.ConfigDirectory, "templates/letter.tex.tt")

	cli := commandline.Cli{}
	err = cli.Run(os.Args, config, current_directory)

	if err != nil {
		os.Exit(1)
	}
}
