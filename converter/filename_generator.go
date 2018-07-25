package converter

import (
	"errors"
	"regexp"
	"strings"
)

type FilenameGenerator struct{}

func NewFilenameGenerator() FilenameGenerator {
	return FilenameGenerator{}
}

func (g *FilenameGenerator) GeneratePdf(input string) (string, error) {
	if input == "" {
		return "", errors.New("Empty input is not allowed to generate file name for PDF file")
	}

	re := regexp.MustCompile("\\.tex")
	escapedString := re.ReplaceAllLiteralString(input, ".pdf")

	return escapedString, nil
}

func (g *FilenameGenerator) Generate(input string) (string, error) {
	if input == "" {
		return "", errors.New("Empty input is not allowed to generate file name for tex file")
	}

	escapedString := strings.ToLower(input)
	re := regexp.MustCompile("[[:blank:]]")
	escapedString = re.ReplaceAllLiteralString(escapedString, "-")
	re = regexp.MustCompile("[^a-z0-9]")
	escapedString = re.ReplaceAllLiteralString(escapedString, "")

	return escapedString, nil
}
