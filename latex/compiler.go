package latex

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fedux-org/letter-generator-go/converter"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Compiler struct {
	Cmd              []string
	Output           string
	WorkingDirectory string
}

func NewCompiler() Compiler {
	return Compiler{}
}

func (c *Compiler) Compile(texFile converter.TexFile) (converter.PdfFile, error) {
	pdfFile := converter.PdfFile{}
	pdfFile.Path = texFile.NameForPdf()

	cmd := exec.Command("pdflatex", "-halt-on-error", texFile.Path, pdfFile.Path)
	cmd.Dir = filepath.Dir(texFile.Path)

	log.WithFields(log.Fields{
		"command":           strings.Join(cmd.Args, " "),
		"working_directory": cmd.Dir,
	}).Debug("Run latex command")

	output, err := cmd.Output()

	var error_messages []string
	error_messages = append(error_messages, string(output))

	if err != nil {
		error_messages = append(error_messages, err.Error())
	}

	if err != nil {
		msg := fmt.Sprintf("exec: Failed to run command \"%s\" in working directory \"%s\".", strings.Join(cmd.Args, " "), cmd.Dir)

		for _, m := range error_messages {
			msg = msg + "\n" + m
		}

		return converter.PdfFile{}, errors.New(msg)
	}

	return pdfFile, nil
}
