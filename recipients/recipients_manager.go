package recipients

import (
	"encoding/json"
	"io/ioutil"
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
	err = json.Unmarshal(data, &recipients)

	if err != nil {
		return err
	}

	m.Recipients = recipients

	return nil
}
