package api

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/fedux-org/letter-generator-go/assets"
	"github.com/fedux-org/letter-generator-go/converter"
	"github.com/fedux-org/letter-generator-go/latex"
	"github.com/fedux-org/letter-generator-go/letter"
	lgos "github.com/fedux-org/letter-generator-go/os"
	log "github.com/sirupsen/logrus"
)

type Project struct {
	letters  []letter.Letter
	template converter.Template
	assets   []assets.Asset
	outDir   string
	workDir  string
}

func NewProject(letters []letter.Letter, template converter.Template, assets []assets.Asset, outDir string) Project {
	p := Project{letters: letters, template: template, assets: assets, outDir: outDir}
	p.SetupWorkDir()

	return p
}

func (p *Project) SetupWorkDir() error {
	dir, err := ioutil.TempDir("", "lg")
	if err != nil {
		return errors.Wrap(err, "create temporary work dir")
	}

	p.workDir = dir

	return nil
}

func (p *Project) Build() error {
	var movableAssets []MovableFile = make([]MovableFile, len(p.assets))
	for i, d := range p.assets {
		movableAssets[i] = &d
	}

	texFiles := generateTexFiles(p.template, p.letters)

	for _, f := range texFiles {
		moveFilesToDir(movableAssets, f.Dir)
	}

	pdfFiles := compileTexFilesIntoPdf(texFiles)

	var movablePdfFiles []MovableFile = make([]MovableFile, len(pdfFiles))
	for i, d := range pdfFiles {
		movablePdfFiles[i] = &d
	}
	moveFilesToDir(movablePdfFiles, p.outDir)

	err := os.RemoveAll(p.workDir)
	if err != nil {
		return errors.Wrap(err, "remove work dir")
	}

	return nil
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

func compileTexFilesIntoPdf(texFiles []converter.TexFile) []converter.PdfFile {
	compiler := latex.NewCompiler()
	var pdfFiles []converter.PdfFile

	for _, f := range texFiles {
		pdfFile, err := compiler.Compile(f)

		if err != nil {
			log.WithFields(log.Fields{
				"input_file":  f.Path,
				"output_file": pdfFile.Path,
			}).Info("Compiling tex file")

			continue
		}

		log.WithFields(log.Fields{
			"input_file":  f.Path,
			"output_file": pdfFile.Path,
		}).Info("Compiling tex file")

		pdfFiles = append(pdfFiles, pdfFile)
	}

	return pdfFiles
}

func createOutputDirectory() (string, error) {
	cwd, err := os.Getwd()
	dir := filepath.Join(cwd, "letters")
	err = os.MkdirAll(dir, 0755)

	if err != nil {
		return "", errors.Wrap(err, "create output directory")
	}

	return cwd, nil
}

func moveFilesToDir(files []MovableFile, dir string) {
	for _, f := range files {
		filename := filepath.Base(f.GetPath())
		newPath := filepath.Join(dir, filename)

		err := lgos.Copy(f.GetPath(), newPath)

		if err != nil {
			log.WithFields(log.Fields{
				"msg":         err.Error(),
				"status":      "failure",
				"source":      f.GetPath(),
				"destination": newPath,
			}).Error("Moving generated pdf file")

			continue
		}

		log.WithFields(log.Fields{
			"output_file": newPath,
		}).Info("Creating letter")
	}
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
