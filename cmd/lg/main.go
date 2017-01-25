package main

import (
	"fmt"
	"github.com/maxmeyer/letter-generator-go/converter"
	"github.com/maxmeyer/letter-generator-go/latex"
	"github.com/maxmeyer/letter-generator-go/letter"
	"github.com/maxmeyer/letter-generator-go/metadata"
	lgos "github.com/maxmeyer/letter-generator-go/os"
	"github.com/maxmeyer/letter-generator-go/recipients"
	"github.com/maxmeyer/letter-generator-go/sender"
	log "github.com/sirupsen/logrus"
	"log"
	_ "net/url"
	"os"
	"path/filepath"
	//"regexp"
	//"strings"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr, could also be a file.
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	metadata := metadata.Metadata{}
	metadata.Read()

	sender := sender.Sender{}
	sender.Read()

	recipient_manager := recipients.RecipientManager{}
	recipient_manager.Read()

	var letters []letter.Letter

	for _, r := range recipient_manager.Recipients {
		lt := letter.New(sender, r, metadata)
		letters = append(letters, lt)
	}

	template := converter.Template{}
	template.Read()

	template_converter := converter.NewConverter()

	var tex_files []converter.TexFile

	for _, l := range letters {
		tex_file, err := template_converter.Transform(l, template)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tex_files = append(tex_files, tex_file)
	}

	compiler := latex.NewCompiler()
	var pdf_files []converter.PdfFile

	for _, f := range tex_files {
		fmt.Println(fmt.Sprintf("Compiling tex file \"%s\".", f.Path))

		pdf_file, err := compiler.Compile(f)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		pdf_files = append(pdf_files, pdf_file)

		fmt.Println(fmt.Sprintf("Generatd pdf file \"%s\".", pdf_file.Path))
	}

	current_working_directory, err := os.Getwd()
	output_directory := filepath.Join(current_working_directory, "letters")
	err = os.MkdirAll(output_directory, 0755)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, f := range pdf_files {
		filename := filepath.Base(f.Path)
		new_path := filepath.Join(output_directory, filename)

		err = lgos.Copy(f.Path, new_path)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(fmt.Sprintf("Moved pdf file from \"%s\" to \"%s\".", f.Path, new_path))
		f.Path = new_path
	}
}
