package metadata

import (
	"encoding/json"
	"io/ioutil"
)

type Metadata struct {
	Subject        string `json:"subject"`
	Signature      string `json:"signature"`
	Opening        string `json:"opening"`
	Closing        string `json:"closing"`
	HasAttachments bool   `json:"has_attachments"`
	HasPs          bool   `json:"has_ps"`
}

func (m *Metadata) Read() error {
	data, err := ioutil.ReadFile("config/metadata.json")

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &m)

	if err != nil {
		return err
	}

	return nil
}
