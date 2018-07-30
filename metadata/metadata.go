package metadata

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Metadata struct {
	Subject        string `yaml:"subject"`
	Signature      string `yaml:"signature"`
	Opening        string `yaml:"opening"`
	Closing        string `yaml:"closing"`
	HasAttachments bool   `yaml:"has_attachments"`
	HasPs          bool   `yaml:"has_ps"`
}

func (m *Metadata) Read(path string) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &m)

	if err != nil {
		return err
	}

	return nil
}
