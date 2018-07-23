package api

import "fmt"

type fileReadError struct {
	path  string
	cause error
}

func (e *fileReadError) Error() string {
	return fmt.Sprintf("Reading file \"%s\" failed %s", e.path, e.cause.Error())
}
