package test

import (
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
)

func Cd(t tester, name string, runTests func()) {
	t.Helper()

	cwd, err := os.Getwd()

	if err != nil {
		t.Fatalf("Capture current working directory failed: %v", err)
	}

	oldpath := cwd
	newpath := ExpandPath(t, name)
	os.Chdir(newpath)
	runTests()
	os.Chdir(oldpath)
}

func CreateEmptyFile(t tester, name string) {
	CreateFile(t, name, "")
}

func CreateFile(t tester, name string, content string) {
	t.Helper()
	var perms os.FileMode

	path := ExpandPath(t, name)
	perms = 0644

	logrus.WithFields(logrus.Fields{
		"path":  path,
		"perms": perms,
	}).Debug("Created file")

	err := ioutil.WriteFile(path, []byte(content), perms)

	if err != nil {
		t.Fatalf("Write path %s failed: %v", path, err)
	}
}

func IsFile(t testHelper, name string) bool {
	t.Helper()
	path := ExpandPath(t, name)
	_, err := os.Stat(path)

	if !os.IsNotExist(err) {
		return true
	}

	return false
}
