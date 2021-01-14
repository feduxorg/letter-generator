package test_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/feduxorg/letter-generator-go/test"
	"github.com/stretchr/testify/assert"
)

func TestWorkingDirectory(t *testing.T) {
	cwd, _ := os.Getwd()
	path := filepath.Join(cwd, "tmp", "aruba")

	assert.Equal(t, path, test.WorkingDirectory())
}

func TestExpandPath(t *testing.T) {
	cwd, _ := os.Getwd()
	name := "file.txt"
	path := filepath.Join(cwd, "tmp", "aruba", name)

	//fmt.Println(test.ExpandPath(t, name))

	assert.Equal(t, path, test.ExpandPath(t, name))
}

func TestSetup(t *testing.T) {
	test.Setup(func() {
		cwd, _ := os.Getwd()
		path := filepath.Join(cwd, "tmp", "aruba")

		_, err := os.Stat(path)

		assert.Equal(t, false, os.IsNotExist(err))
	})

}
