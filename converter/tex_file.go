package converter

import (
	"io/ioutil"
	"path/filepath"
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

func (t *TexFile) NameForPdf() string {
	return filepath.Join(t.Dir, t.Name+".pdf")
}
