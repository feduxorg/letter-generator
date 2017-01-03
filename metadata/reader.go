package metadata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func (m *Metadata) Read() *Metadata {
	data, err := ioutil.ReadFile("config/metadata.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &m)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return m
}
