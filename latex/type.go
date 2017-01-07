package latex

import (
	"errors"
	"fmt"
	"github.com/maxmeyer/letter-generator-go/letter"
	"os/exec"
	"path/filepath"
)

type Compiler struct {
	CmdString        string
	Output           string
	WorkingDirectory string
}

func (c *Compiler) Compile(file letter.Letter) error {
	c.CmdString = fmt.Sprintf("pdflatex -halt-on-error %s %s", file.TexPath, file.PdfPath)
	c.WorkingDirectory = filepath.Dir(file.TexPath)

	cmd := exec.Command(c.CmdString)
	cmd.Dir = c.WorkingDirectory
	output, err := cmd.Output()

	if err != nil {
		msg := fmt.Sprintf("exec: Failed to run command \"%s\" in working directory \"%s\".\n\n%s", c.CmdString, c.WorkingDirectory, string(output))
		return errors.New(msg)
	}

	return nil
}
