package main

import (
	"fmt"
	"github.com/maxmeyer/letter-generator-go/converter"
	"github.com/maxmeyer/letter-generator-go/metadata"
	"github.com/maxmeyer/letter-generator-go/recipients"
	"github.com/maxmeyer/letter-generator-go/sender"
	"io/ioutil"
	"os"
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

	converter := converter.TemplateConverter{}

	for _, r := range recipient_manager.Recipients {
		output_file, err := ioutil.TempFile("", "letter_template_XXX.tex")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		converter.Transform(
			&metadata,
			&r,
			&sender,
			template,
			output_file,
		)
	}
}
