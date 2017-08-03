package commandline

import (
	"errors"
	_ "path/filepath"
	"strings"

	"github.com/fedux-org/letter-generator-go/letter_generator"
	"github.com/libgit2/git2go"
	log "github.com/sirupsen/logrus"
)

type Initializer struct{}

func (i *Initializer) Init(dir string, config letter_generator.Config) error {
	//dir = filepath.Join(dir, config.ConfigDirectory)

	log.WithFields(log.Fields{
		"known_sources": strings.Join(config.RemoteSources, ", "),
		"destination":   dir,
	}).Debug("Starting to clone config repository")

	var remote_source_cloned bool
	remote_source_cloned = false

	for _, s := range config.RemoteSources {
		if remote_source_cloned == true {
			break
		}

		repo, err := git.Clone(s, dir, &git.CloneOptions{})

		if err != nil {
			log.WithFields(log.Fields{
				"source":      s,
				"destination": dir,
				"msg":         err.Error(),
				"status":      "failure",
			}).Debug("Cloning config repository")

			continue
		}

		remote := "origin"
		repo.Remotes.Delete(remote)

		if err != nil {
			log.WithFields(log.Fields{
				"repository": repo,
				"remote":     remote,
				"msg":        err.Error(),
				"status":     "failure",
			}).Warn("Removing remote failed")
		}

		remote_source_cloned = true

		log.WithFields(log.Fields{
			"source":      s,
			"destination": dir,
			"status":      "success",
		}).Debug("Cloning config repository")
	}

	if remote_source_cloned == false {

		log.WithFields(log.Fields{
			"msg":    "No valid remote sources found",
			"status": "failure",
		}).Fatal("Cloning config repository")

		return errors.New("No valid remote sources found")
	}

	return nil
}
