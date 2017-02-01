package os

import (
	"os/user"
)

func HomeDirectory() (string, error) {
	usr, err := user.Current()

	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}
