package main

import (
	"github.com/maxmeyer/letter-generator-go/converter"
	"github.com/maxmeyer/letter-generator-go/latex"
	"github.com/maxmeyer/letter-generator-go/letter"
	"github.com/maxmeyer/letter-generator-go/metadata"
	lgos "github.com/maxmeyer/letter-generator-go/os"
	"github.com/maxmeyer/letter-generator-go/recipients"
	"github.com/maxmeyer/letter-generator-go/sender"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	_ "net/url"
	"os"
	"path/filepath"
	//"regexp"
	//"strings"
)

var (
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr, could also be a file.
	// log.SetOutput(os.Stdout)
}

func main() {
	kingpin.Parse()

	if *verbose == true {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	metadata := metadata.Metadata{}
	err := metadata.Read()

	if err != nil {
		log.WithFields(log.Fields{
			"msg":    err.Error(),
			"status": "failure",
		}).Fatal("Reading letter metadata")

		os.Exit(1)
	}

	log.WithFields(log.Fields{
		"status": "success",
	}).Debug("Reading metadata")

	sender := sender.Sender{}
	err = sender.Read()

	if err != nil {
		log.WithFields(log.Fields{
			"msg":    err.Error(),
			"status": "failure",
		}).Fatal("Reading sender")

		os.Exit(1)
	}

	log.WithFields(log.Fields{
		"status": "success",
	}).Debug("Reading sender")

	recipient_manager := recipients.RecipientManager{}
	err = recipient_manager.Read()

	log.WithFields(log.Fields{
		"valid":  len(recipient_manager.Recipients),
		"status": "success",
	}).Info("Read recipients")

	var letters []letter.Letter

	for _, r := range recipient_manager.Recipients {
		lt := letter.New(sender, r, metadata)
		letters = append(letters, lt)
	}

	log.WithFields(log.Fields{
		"status": "success",
	}).Debug("Create letter instances")

	template := converter.Template{}
	template.Read()

	if err != nil {
		log.WithFields(log.Fields{
			"msg":    err.Error(),
			"status": "failure",
		}).Fatal("Reading letter template")

		os.Exit(1)
	}

	log.WithFields(log.Fields{
		"status": "success",
	}).Debug("Reading letter template")

	template_converter := converter.NewConverter()

	var tex_files []converter.TexFile

	for _, l := range letters {
		tex_file, err := template_converter.Transform(l, template)

		if err != nil {
			log.WithFields(log.Fields{
				"msg":    err.Error(),
				"status": "failure",
			}).Fatal("Create tex files")

			os.Exit(1)
		}

		log.WithFields(log.Fields{
			"status": "success",
			"path":   tex_file.Path,
		}).Debug("Creating tex file from template")

		tex_files = append(tex_files, tex_file)
	}

	compiler := latex.NewCompiler()
	var pdf_files []converter.PdfFile

	for _, f := range tex_files {
		pdf_file, err := compiler.Compile(f)

		if err != nil {
			log.WithFields(log.Fields{
				"msg":    err.Error(),
				"status": "failure",
			}).Fatal("Compiling tex files")

			os.Exit(1)
		}

		log.WithFields(log.Fields{
			"input_file":  f.Path,
			"output_file": pdf_file.Path,
			"status":      "success",
		}).Info("Compiling tex file")

		pdf_files = append(pdf_files, pdf_file)
	}

	current_working_directory, err := os.Getwd()
	output_directory := filepath.Join(current_working_directory, "letters")
	err = os.MkdirAll(output_directory, 0755)

	if err != nil {
		log.WithFields(log.Fields{
			"path":   output_directory,
			"status": "failure",
		}).Fatal("Generating output directory")
		os.Exit(1)
	}

	for _, f := range pdf_files {
		filename := filepath.Base(f.Path)
		new_path := filepath.Join(output_directory, filename)

		err = lgos.Copy(f.Path, new_path)

		if err != nil {

			log.WithFields(log.Fields{
				"msg":         err.Error(),
				"status":      "failure",
				"source":      f.Path,
				"destination": new_path,
			}).Fatal("Moving generatd pdf file")

			os.Exit(1)
		}

		log.WithFields(log.Fields{
			"output_file": f.Path,
			"status":      "success",
		}).Info("Creating letter")

		f.Path = new_path
	}
}
