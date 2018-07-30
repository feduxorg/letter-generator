package sender

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Sender struct {
	Name   string `yaml:"name"`
	Street string `yaml:"street"`
	City   string `yaml:"city"`
	Phone  string `yaml:"Phone"`
	Mail   string `yaml:"Mail"`
}

func (s *Sender) Read(path string) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &s)

	if err != nil {
		return err
	}

	return nil
}
