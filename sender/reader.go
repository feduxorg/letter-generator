package sender

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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
