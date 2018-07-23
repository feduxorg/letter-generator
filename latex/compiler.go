package latex

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/fedux-org/letter-generator-go/converter"
)

type Compiler struct {
	CmdString        string
	Output           string
	WorkingDirectory string
}

func NewCompiler() Compiler {
	return Compiler{}
}

func (c *Compiler) Compile(texFile converter.TexFile) (converter.PdfFile, error) {
	filename_generator := converter.NewFilenameGenerator()

	path, err := filename_generator.GeneratePdf(texFile.Path)

	if err != nil {
		return converter.PdfFile{}, err
	}

	pdfFile := converter.PdfFile{}
	pdfFile.Path = path

	c.CmdString = "pdflatex " + " -halt-on-error " + " " + texFile.Path + " " + pdfFile.Path
	c.WorkingDirectory = filepath.Dir(texFile.Path)

	cmd := exec.Command("pdflatex", "-halt-on-error", texFile.Path, pdfFile.Path)
	cmd.Dir = c.WorkingDirectory
	output, err := cmd.Output()

	var error_messages []string
	error_messages = append(error_messages, string(output))

	if err != nil {
		error_messages = append(error_messages, err.Error())
	}

	if err != nil {
		msg := fmt.Sprintf("exec: Failed to run command \"%s\" in working directory \"%s\".", c.CmdString, c.WorkingDirectory)

		for _, m := range error_messages {
			msg = msg + "\n" + m
		}

		return converter.PdfFile{}, errors.New(msg)
	}

	return pdfFile, nil
}
