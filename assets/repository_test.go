pGOOS=windows GOARCH=amd64 GOROOT_BOOTSTRAP=/usr/lib/go ./make.bash
# [...]
# => Installed Go for windows/amd64 in /root/.local/share/go-source
# => Installed commands in /root/.local/share/go-source/bin

GOOS=darwin GOARCH=amd64 GOROOT_BOOTSTRAP=/usr/lib/go ./make.bashackage assets_test

import (
	"os"
	"testing"

	"github.com/feduxorg/letter-generator-go/assets"
	"github.com/feduxorg/letter-generator-go/test"
	"github.com/stretchr/testify/assert"
)

var cwd string

func TestMain(m *testing.M) {
	var retCode int

	test.Setup(func() {
		retCode = m.Run()
	})

	os.Exit(retCode)
}

func TestAddAsset(t *testing.T) {
	repo := assets.Repository{}
	repo.AddAsset(test.ExpandPath(t, "file1.txt"))
	repo.AddAsset(test.ExpandPath(t, "file2.txt"))

	assert.Equal(t, 2, len(repo.KnownAssets()))
}

func TestInit(t *testing.T) {
	test.CreateEmptyFile(t, "file1.txt")
	test.CreateEmptyFile(t, "file2.txt")

	rootDir := test.ExpandPath(t, ".")
	repo := assets.Repository{Path: rootDir}
	repo.Init()

	assert.Equal(t, 2, len(repo.KnownAssets()))
}
