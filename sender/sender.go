package sender

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Sender struct {
	Name   string `json:"name"`
	Street string `json:"street"`
	City   string `json:"city"`
	Phone  string `json:"Phone"`
	Mail   string `json:"Mail"`
}

func (s *Sender) Read() *Sender {
	data, err := ioutil.ReadFile("config/from.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &s)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return s
}
