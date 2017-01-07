package converter

import (
	_ "fmt"
	"github.com/maxmeyer/letter-generator-go/metadata"
	"github.com/maxmeyer/letter-generator-go/recipients"
	"github.com/maxmeyer/letter-generator-go/sender"
	"io"
	_ "os"
	gotmpl "text/template"
)

type TemplateConverter struct{}

type Context struct {
	Metadata  *metadata.Metadata
	Recipient *recipients.Recipient
	Sender    *sender.Sender
}

func (c *TemplateConverter) Transform(
	metadata *metadata.Metadata,
	recipient *recipients.Recipient,
	sender *sender.Sender,
	template Template,
	output_file io.Writer,
) {

	tmpl, err := gotmpl.New("letter_template").Parse(template.Content)

	if err != nil {
		panic(err)
	}

	context := Context{
		Metadata:  metadata,
		Recipient: recipient,
		Sender:    sender,
	}

	err = tmpl.Execute(output_file, context)

	if err != nil {
		panic(err)
	}

	// fmt.Println(output_file)
}
