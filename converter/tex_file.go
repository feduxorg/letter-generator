package converter

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
)

type TexFile struct {
	Path string
	Name string
	Dir  string
}

func NewTexFile(fileName string) (TexFile, error) {
	outputDir, err := ioutil.TempDir("", "letter_template_")

	if err != nil {
		return TexFile{}, err
	}

	outputDir, err = filepath.Abs(outputDir)

	if err != nil {
		return TexFile{}, err
	}

	outputPath := filepath.Join(outputDir, fileName+".tex")

	if err != nil {
		return TexFile{}, err
	}

	texFile := TexFile{
		Dir:  outputDir,
		Path: outputPath,
		Name: fileName,
	}

	return texFile, nil
}

func (f *TexFile) NameForPdf() string {
	return filepath.Join(f.Dir, f.Name+".pdf")
}

func (f *TexFile) Destroy() error {
	log.WithField("directory", f.Dir).Debug("Remove build dir")

	if f.Dir == "" {
		log.WithField("result", false).Debug("Verify directory is set")
		return nil
	}

	if _, err := os.Stat(f.Dir); os.IsNotExist(err) {
		log.WithFields(log.Fields{
			"path":   f.Dir,
			"result": false,
		}).Debug("Verify directory exists")

		return nil
	}

	return os.RemoveAll(f.Dir)
}
