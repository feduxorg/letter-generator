package converter

import (
	"io/ioutil"
)

type Template struct {
	Content string
}

func (t *Template) Read(path string) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	t.Content = string(data)

	return nil
}
