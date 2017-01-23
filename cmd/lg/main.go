package main

import (
	"fmt"
	"github.com/maxmeyer/letter-generator-go/converter"
	"github.com/maxmeyer/letter-generator-go/latex"
	"github.com/maxmeyer/letter-generator-go/letter"
	"github.com/maxmeyer/letter-generator-go/metadata"
	"github.com/maxmeyer/letter-generator-go/recipients"
	"github.com/maxmeyer/letter-generator-go/sender"
	_ "net/url"
	"os"
	//"regexp"
	//"strings"
)

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

	for _, f := range tex_files {
		fmt.Println(fmt.Sprintf("Compiling tex file \"%s\".", f.Path))

		pdf_file, err := compiler.Compile(f.Path)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(fmt.Sprintf("Generatd pdf file \"%s\".", pdf_file))
	}
}
