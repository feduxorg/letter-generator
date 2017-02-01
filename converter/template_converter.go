package converter

import (
	"github.com/maxmeyer/letter-generator-go/letter"
	"os"
	gotmpl "text/template"
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

	filename_generator := NewFilenameGenerator()
	file_name, err := filename_generator.GenerateTex(context.Recipient.Name)

	if err != nil {
		return TexFile{}, err
	}

	tex_file, err := NewTexFile(file_name)

	if err != nil {
		return TexFile{}, err
	}

	output_file, err := os.Create(tex_file.Path)

	if err != nil {
		return TexFile{}, err
	}

	err = tmpl.Execute(output_file, context)

	if err != nil {
		return TexFile{}, err
	}

	return tex_file, nil

	// fmt.Println(output_file)
}
