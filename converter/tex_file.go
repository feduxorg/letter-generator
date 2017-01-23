package converter

import (
	"io/ioutil"
	"path/filepath"
)

type TexFile struct {
	Path string
}

func NewTexFile(file_name string) (TexFile, error) {
	output_directory, err := ioutil.TempDir("", "letter_template_")

	if err != nil {
		return TexFile{}, err
	}

	output_directory, err = filepath.Abs(output_directory)

	if err != nil {
		return TexFile{}, err
	}

	output_path := filepath.Join(output_directory, file_name)

	if err != nil {
		return TexFile{}, err
	}

	tex_file := TexFile{
		Path: output_path,
	}

	return tex_file, nil
}

//new_filename := fmt.Sprintf("%s.pdf", url.QueryEscape())
//escaped_string = strings.Replace(lt.TexPath, old_filename, new_filename, -1)
