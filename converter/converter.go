package converter

import (
	"fmt"
	"github.com/maxmeyer/letter-generator-go/metadata"
	"github.com/maxmeyer/letter-generator-go/recipients"
	"github.com/maxmeyer/letter-generator-go/sender"
	"io"
	"os"
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

	tmpl, err := gotmpl.New("test").Parse(template.Content)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	context := Context{
		Metadata:  metadata,
		Recipient: recipient,
		Sender:    sender,
	}

	err = tmpl.Execute(output_file, context)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
