package test

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var verboseLogging bool

func init() {
	flag.BoolVar(&verboseLogging, "debug", false, "activate logging for tests")
	flag.BoolVar(&verboseLogging, "d", false, "activate logging for tests")
}

func ExpandPath(t testHelper, name string) string {
	t.Helper()
	path := filepath.Join(WorkingDirectory(), name)

	logrus.WithFields(logrus.Fields{
		"path":              path,
		"working_directory": WorkingDirectory(),
	}).Debug("Expanded path")

	return path
}

func Setup(runTests func()) {
	flag.Parse()

	if verboseLogging == true {
		logrus.SetLevel(logrus.DebugLevel)
	}

	os.RemoveAll(WorkingDirectory())
	os.MkdirAll(WorkingDirectory(), 0755)

	logrus.WithFields(logrus.Fields{
		"path": WorkingDirectory(),
	}).Debug("Set up test environment")

	runTests()
}

func WorkingDirectory() string {
	cwd, err := os.Getwd()

	if err != nil {
		cwd = "/tmp"
	}

	path := filepath.Join(cwd, "tmp", "aruba")

	logrus.WithFields(logrus.Fields{
		"path": path,
	}).Debug("Got working directory")

	return path
}
