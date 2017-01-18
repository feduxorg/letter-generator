package letter

import (
	"fmt"
	"github.com/maxmeyer/letter-generator-go/recipients"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Letter struct {
	Recipient       recipients.Recipient
	OutputDirectory string
	TexFile         io.Writer
	TexPath         string
	PdfPath         string
}

func New(recipient recipients.Recipient) Letter {
	letter := Letter{}
	letter.Recipient = recipient

	letter.CreateOutputDirectory()
	letter.CreateTexFile()

	return letter
}

func (l *Letter) GeneratePdfPath() {
	escaped_string := strings.ToLower(r.Name)
	//old_filename := "index.tex"
	re := regexp.MustCompile("[[:blank:]]")
	escaped_string = re.ReplaceAllLiteralString(escaped_string, "-")
	re = regexp.MustCompile("[^a-z0-9]")
	escaped_string = re.ReplaceAllLiteralString(escaped_string, "")
}

func (l *Letter) CreateTexFile() {
	output_path := filepath.Join(l.OutputDirectory, "index.tex")
	output_file, err := os.Create(output_path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l.TexPath = output_path
	l.TexFile = output_file
}

func (l *Letter) CreateOutputDirectory() {
	output_directory, err := ioutil.TempDir("", "letter_template_")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output_directory, err = filepath.Abs(output_directory)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l.OutputDirectory = output_directory
}
