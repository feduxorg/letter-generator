package recipients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func (m *RecipientManager) Read() *RecipientManager {
	data, err := ioutil.ReadFile("config/to.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var recipients []Recipient
	err = json.Unmarshal(data, &recipients)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m.Recipients = recipients

	return m
}
