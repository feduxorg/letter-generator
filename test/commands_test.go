package test_test

import (
	"testing"
	"time"

	"os/exec"

	"github.com/feduxorg/letter-generator/test"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestRunCommand(t *testing.T) {
	message := "Hello, Aruba!"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := test.Command(exec.CommandContext(ctx, "echo", "-n", message))
	test.RunCommand(t, cmd)

	assert.Equal(t, message, cmd.ReadStdout())
}
