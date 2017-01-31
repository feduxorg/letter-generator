package commandline

import (
	"errors"
	"github.com/libgit2/git2go"
	"github.com/maxmeyer/letter-generator-go/letter_generator"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Initializer struct{}

func (i *Initializer) Init(dir string, config letter_generator.Config) error {
	log.WithFields(log.Fields{
		"sources":     strings.Join(config.RemoteSources, ", "),
		"destination": dir,
	}).Debug("Cloning config repository")

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
				"result":      repo,
				"status":      "failure",
			}).Debug("Cloning config repository")
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
