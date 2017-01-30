package sender

import (
	"encoding/json"
	"io/ioutil"
)

type Sender struct {
	Name   string `json:"name"`
	Street string `json:"street"`
	City   string `json:"city"`
	Phone  string `json:"Phone"`
	Mail   string `json:"Mail"`
}

func (s *Sender) Read(path string) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &s)

	if err != nil {
		return err
	}

	return nil
}
