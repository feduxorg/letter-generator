package main

import (
	"github.com/maxmeyer/letter-generator-go/commandline"
	"github.com/maxmeyer/letter-generator-go/letter_generator"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
)

var (
	verbose      = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	buildCommand = kingpin.Command("build", "Build letters")
	initCommand  = kingpin.Command("init", "Initialize project directory to build letters")
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr, could also be a file.
	// log.SetOutput(os.Stdout)
}

func main() {
	builder := commandline.LetterBuilder{}

	current_directory, err := os.Getwd()

	if err != nil {
		log.WithFields(log.Fields{
			"msg":    err.Error(),
			"status": "failure",
		}).Fatal("Getting current directory")

		os.Exit(1)
	}

	config_directory := ".lg"

	log.WithFields(log.Fields{
		"path":   current_directory,
		"status": "success",
	}).Debug("Getting current directory")

	config := letter_generator.Config{}
	config.RecipientsFile = filepath.Join(current_directory, config_directory, "source/to.json")
	config.MetadataFile = filepath.Join(current_directory, config_directory, "source/metadata.json")
	config.SenderFile = filepath.Join(current_directory, config_directory, "source/from.json")
	config.TemplateFile = filepath.Join(current_directory, config_directory, "templates/letter.tex.tt")

	if *verbose == true {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	switch kingpin.Parse() {
	case "build":
		builder.Build(config)
	case "init":
	default:
		builder.Build(config)
	}

}
