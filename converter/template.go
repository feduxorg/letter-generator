package converter

import (
	"io/ioutil"
)

type Template struct {
	Content string
	Path    string
}

func (t *Template) Read(path string) error {
	t.Path = path

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	t.Content = string(data)

	return nil
}
