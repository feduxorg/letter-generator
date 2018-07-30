package recipients

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type RecipientManager struct {
	Recipients []Recipient
}

func (m *RecipientManager) Read(path string) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	var recipients []Recipient
	err = yaml.Unmarshal(data, &recipients)

	if err != nil {
		return err
	}

	m.Recipients = recipients

	return nil
}
