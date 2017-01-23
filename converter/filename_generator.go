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
		return "", errors.New("Empty input")
	}

	re := regexp.MustCompile("\\.tex")
	escaped_string := re.ReplaceAllLiteralString(input, ".pdf")

	return escaped_string, nil
}

func (g *FilenameGenerator) GenerateTex(input string) (string, error) {
	if input == "" {
		return "", errors.New("Empty input")
	}

	escaped_string := strings.ToLower(input)
	re := regexp.MustCompile("[[:blank:]]")
	escaped_string = re.ReplaceAllLiteralString(escaped_string, "-")
	re = regexp.MustCompile("[^a-z0-9]")
	escaped_string = re.ReplaceAllLiteralString(escaped_string, "")

	escaped_string = escaped_string + ".tex"

	return escaped_string, nil
}
