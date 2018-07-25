package converter

import (
	"os"
	gotmpl "text/template"

	"github.com/fedux-org/letter-generator-go/letter"
	log "github.com/sirupsen/logrus"
)

type TemplateConverter struct{}

func NewConverter() TemplateConverter {
	return TemplateConverter{}
}

func (c *TemplateConverter) Transform(
	letter letter.Letter,
	template Template,
) (TexFile, error) {

	tmpl, err := gotmpl.New("letter_template").Parse(template.Content)

	if err != nil {
		return TexFile{}, err
	}

	context := TemplateContext{
		Recipient:      &letter.Recipient,
		Sender:         &letter.Sender,
		Subject:        letter.Subject,
		Signature:      letter.Signature,
		Opening:        letter.Opening,
		Closing:        letter.Closing,
		HasAttachments: letter.HasAttachments,
		HasPs:          letter.HasPs,
	}

	nameGen := NewFilenameGenerator()
	fileName, err := nameGen.Generate(context.Recipient.Name)

	if err != nil {
		return TexFile{}, err
	}

	texFile, err := NewTexFile(fileName)

	if err != nil {
		return TexFile{}, err
	}

	outputFile, err := os.Create(texFile.Path)

	log.WithFields(log.Fields{
		"file_name": fileName,
		"path":      outputFile,
	}).Debug("Create new tex file")

	if err != nil {
		return TexFile{}, err
	}

	err = tmpl.Execute(outputFile, context)
	if err != nil {
		return TexFile{}, err
	}

	log.WithFields(log.Fields{
		"file_name": fileName,
		"path":      outputFile,
	}).Debug("Render template into tex file")

	return texFile, nil
}
