package converter

import (
	"os"
	gotmpl "text/template"

	"github.com/feduxorg/letter-generator/letter"
	"github.com/pkg/errors"
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
		"file_name": texFile.Name,
		"path":      texFile.Path,
	}).Debug("Create new tex file")

	if err != nil {
		return TexFile{}, err
	}

	tmpl, err := gotmpl.New(template.Path).Parse(template.Content)

	if err != nil {
		return TexFile{}, errors.Wrap(err, "create template instance")
	}

	err = tmpl.Execute(outputFile, context)
	if err != nil {
		return texFile, errors.Wrap(err, "render template")
	}

	return texFile, nil
}
