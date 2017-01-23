package latex

import (
	"errors"
	"fmt"
	"github.com/maxmeyer/letter-generator-go/converter"
	"os/exec"
	"path/filepath"
)

type Compiler struct {
	CmdString        string
	Output           string
	WorkingDirectory string
}

func NewCompiler() Compiler {
	return Compiler{}
}

func (c *Compiler) Compile(tex_file string) (string, error) {
	filename_generator := converter.NewFilenameGenerator()
	pdf_file, err := filename_generator.GeneratePdf(tex_file)

	if err != nil {
		return "", err
	}

	c.CmdString = "pdflatex " + " -halt-on-error " + " " + tex_file + " " + pdf_file
	c.WorkingDirectory = filepath.Dir(tex_file)

	cmd := exec.Command("pdflatex", "-halt-on-error", tex_file, pdf_file)
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

		return "", errors.New(msg)
	}

	return pdf_file, nil
}
