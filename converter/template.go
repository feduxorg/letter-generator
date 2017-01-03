package converter

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Template struct {
	Content string
}

func (t *Template) Read() *Template {
	data, err := ioutil.ReadFile("templates/letter.tex.tt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t.Content = string(data)

	return t
}
