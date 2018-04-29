package test

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

type Cmd struct {
	exec.Cmd
	stdoutCache *bytes.Buffer
	stderrCache *bytes.Buffer
}

func Command(cmd *exec.Cmd) *Cmd {
	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	return &Cmd{
		Cmd:         *cmd,
		stdoutCache: &stdout,
		stderrCache: &stderr,
	}
}

func (c *Cmd) ReadStdout() string {
	return c.stdoutCache.String()
}

func (c *Cmd) ReadStderr() string {
	return c.stderrCache.String()
}

func RunCommand(t tester, cmd *Cmd) {
	t.Helper()

	var cmdPath string
	var cmdArgs []string

	if len(cmd.Cmd.Args) > 0 {
		cmdPath = cmd.Cmd.Args[0]
		cmdArgs = cmd.Cmd.Args[1:]
	}

	if err := cmd.Run(); err != nil {
		logrus.WithFields(logrus.Fields{
			"cmd":    cmdPath,
			"args":   strings.Join(cmdArgs, " "),
			"stdout": truncate(t, cmd.ReadStdout(), 30),
			"stderr": truncate(t, cmd.ReadStderr(), 30),
			"error":  err.Error(),
		}).Error("Ran command")

		t.Fatalf("%v", err)
	}

	logrus.WithFields(logrus.Fields{
		"cmd":    cmdPath,
		"args":   strings.Join(cmdArgs, " "),
		"stdout": truncate(t, cmd.ReadStdout(), 30),
		"stderr": truncate(t, cmd.ReadStderr(), 30),
	}).Debug("Ran command")
}
