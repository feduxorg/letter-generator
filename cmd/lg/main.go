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
	"regexp"
	"strings"
)

func main() {
	metadata := metadata.Metadata{}
	metadata.Read()

	recipient_manager := recipients.RecipientManager{}
	recipient_manager.Read()

	sender := sender.Sender{}
	sender.Read()

	template := converter.Template{}
	template.Read()

	template_converter := converter.TemplateConverter{}
	compiler := latex.Compiler{}

	var letters []letter.Letter

	for _, r := range recipient_manager.Recipients {
		lt := letter.New(r)
		letters = append(letters, lt)

		fmt.Println(escaped_string)

		os.Exit(1)

		//new_filename := fmt.Sprintf("%s.pdf", url.QueryEscape())
		//escaped_string = strings.Replace(lt.TexPath, old_filename, new_filename, -1)

		fmt.Println(lt.PdfPath)

		//lt.PdfPath = pdf_path_converter.Convert(lt.TexPath)

		template_converter.Transform(
			&metadata,
			&r,
			&sender,
			template,
			lt.TexFile,
		)
	}

	for _, l := range letters {
		err := compiler.Compile(l)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}
}
