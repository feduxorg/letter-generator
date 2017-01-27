package recipients

import (
	"encoding/json"
	"io/ioutil"
)

type RecipientManager struct {
	Recipients []Recipient
}

func (m *RecipientManager) Read() error {
	data, err := ioutil.ReadFile("config/to.json")

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
