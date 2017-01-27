package converter

import (
	"io/ioutil"
)

type Template struct {
	Content string
}

func (t *Template) Read() error {
	data, err := ioutil.ReadFile("templates/letter.tex.tt")

	if err != nil {
		return err
	}

	t.Content = string(data)

	return nil
}
