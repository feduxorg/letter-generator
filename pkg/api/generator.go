package api

import (
	"github.com/pkg/errors"
	_ "net/url"
	"os"
	"path/filepath"

	"github.com/fedux-org/letter-generator-go/assets"
	"github.com/fedux-org/letter-generator-go/converter"
	"github.com/fedux-org/letter-generator-go/latex"
	"github.com/fedux-org/letter-generator-go/letter"
	"github.com/fedux-org/letter-generator-go/letter_generator"
	"github.com/fedux-org/letter-generator-go/metadata"
	lgos "github.com/fedux-org/letter-generator-go/os"
	"github.com/fedux-org/letter-generator-go/recipients"
	"github.com/fedux-org/letter-generator-go/sender"
	log "github.com/sirupsen/logrus"
)

type LetterBuilder struct{}

func (lc *LetterBuilder) Build(config letter_generator.Config) error {
	metadata, err := readMetadata(config.MetadataFile)
	if err != nil {
		return err
	}

	sender, err := readSender(config.SenderFile)
	if err != nil {
		return err
	}

	recipientManager, err := readRecipients(config.RecipientsFile)
	if err != nil {
		return err
	}

	template, err := readLetterTemplate(config.TemplateFile)
	if err != nil {
		return err
	}

	assetRepo := assets.NewRepository(config.AssetsDirectory)

	if err := assetRepo.Init(); err != nil {
		return errors.Wrap(err, "init repo")
	}

	log.WithFields(log.Fields{
		"status":      "success",
		"len(assets)": len(assetRepo.KnownAssets()),
	}).Debug("Initializing repository for assets")

	letters := generateLetterInstances(sender, metadata, recipientManager.Recipients)
	texFiles := generateTexFiles(template, letters)

	compiler := latex.NewCompiler()
	var pdf_files []converter.PdfFile

	for _, f := range texFiles {
		pdf_file, err := compiler.Compile(f)

		if err != nil {
			log.WithFields(log.Fields{
				"msg":    err.Error(),
				"status": "failure",
			}).Fatal("Compiling tex files")

			return err
		}

		log.WithFields(log.Fields{
			"input_file":  f.Path,
			"output_file": pdf_file.Path,
			"status":      "success",
		}).Info("Compiling tex file")

		pdf_files = append(pdf_files, pdf_file)
	}

	current_working_directory, err := os.Getwd()
	output_directory := filepath.Join(current_working_directory, "letters")
	err = os.MkdirAll(output_directory, 0755)

	if err != nil {
		log.WithFields(log.Fields{
			"path":   output_directory,
			"status": "failure",
		}).Fatal("Generating output directory")

		return err
	}

	for _, f := range pdf_files {
		filename := filepath.Base(f.Path)
		new_path := filepath.Join(output_directory, filename)

		err = lgos.Copy(f.Path, new_path)

		if err != nil {

			log.WithFields(log.Fields{
				"msg":         err.Error(),
				"status":      "failure",
				"source":      f.Path,
				"destination": new_path,
			}).Fatal("Moving generaed pdf file")

			return err
		}

		log.WithFields(log.Fields{
			"output_file": new_path,
			"status":      "success",
		}).Info("Creating letter")

		f.Path = new_path
	}

	return nil
}

func readMetadata(srcFile string) (metadata.Metadata, error) {
	m := metadata.Metadata{}
	err := m.Read(srcFile)

	if err != nil {
		return metadata.Metadata{}, errors.Wrap(err, "read metadata")
	}

	log.WithField("file", srcFile).Debug("Reading metadata")

	return m, nil
}

func readSender(srcFile string) (sender.Sender, error) {
	s := sender.Sender{}
	err := s.Read(srcFile)

	if err != nil {
		return sender.Sender{}, errors.Wrap(err, "read sender information")
	}

	log.WithField("file", srcFile).Debug("Reading sender")

	return s, nil
}

func readRecipients(srcFile string) (recipients.RecipientManager, error) {
	recipientManager := recipients.RecipientManager{}

	err := recipientManager.Read(srcFile)
	if err != nil {
		return recipients.RecipientManager{}, errors.Wrap(err, "read recipients files")
	}

	logInfo := log.WithFields(log.Fields{
		"count(valid_recipients)": len(recipientManager.Recipients),
	})

	logInfo.Info("Reading recipients")
	logInfo.WithField("file", srcFile).Debug("Reading recipients")

	return recipientManager, nil
}

func generateLetterInstances(sender sender.Sender, metadata metadata.Metadata, recipients []recipients.Recipient) []letter.Letter {
	var letters []letter.Letter

	for _, r := range recipients {
		lt := letter.New(sender, r, metadata)
		letters = append(letters, lt)
	}

	log.WithFields(log.Fields{
		"count(letters)": len(letters),
	}).Debug("Creating letter instances")

	return letters
}

func readLetterTemplate(srcFile string) (converter.Template, error) {
	template := converter.Template{}

	err := template.Read(srcFile)
	if err != nil {
		return converter.Template{}, errors.Wrap(err, "read template")
	}

	log.WithFields(log.Fields{
		"file": srcFile,
	}).Debug("Reading letter template")

	return template, nil
}

func renderTemplate(l letter.Letter, t converter.Template) (converter.TexFile, error) {
	templateConverter := converter.NewConverter()
	texFile, err := templateConverter.Transform(l, t)

	if err != nil {
		return converter.TexFile{}, errors.Wrap(err, "render template into tex file")
	}

	log.WithFields(log.Fields{
		"path(tex_file)": texFile.Path,
		"path(template)": t.Path,
	}).Debug("Creating tex file from template")

	return texFile, nil

}

func generateTexFiles(template converter.Template, letters []letter.Letter) []converter.TexFile {
	var texFiles []converter.TexFile

	for _, l := range letters {
		texFile, err := renderTemplate(l, template)

		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"letter":         l,
				"path(template)": template.Path,
			}).Error("Render letter into template")

			continue
		}

		texFiles = append(texFiles, texFile)
	}

	return texFiles
}
