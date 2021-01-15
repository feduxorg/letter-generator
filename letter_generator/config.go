package letter_generator

import (
	"fmt"
)

// The current build version.
const AppVersionNumber = "dev"

// SHA-value of git commit
const CommitHash = ""

// Date of build
const BuildDate = ""

type Config struct {
	ConfigDirectory string `yaml:config_directory`
	TemplateSource  string `yaml:template_source`
	RecipientsFile  string `yaml:recipients_file`
	MetadataFile    string `yaml:metadata_file`
	SenderFile      string `yaml:sender_file`
	TemplateFile    string `yaml:template_file`
	AssetsDirectory string `yaml:assets_directory`
}

func (c *Config) ToString() []string {
	result := []string{}
	result = append(result, fmt.Sprintf("%20s | %-30s", "Option", "Value"))
	result = append(result, fmt.Sprintf("%20s | %-30s", "Template Source", c.TemplateSource))
	result = append(result, fmt.Sprintf("%20s | %-30s", "Recipients File", c.RecipientsFile))
	result = append(result, fmt.Sprintf("%20s | %-30s", "Metadata File", c.MetadataFile))
	result = append(result, fmt.Sprintf("%20s | %-30s", "Sender File", c.SenderFile))
	result = append(result, fmt.Sprintf("%20s | %-30s", "Template File", c.TemplateFile))
	result = append(result, fmt.Sprintf("%20s | %-30s", "Assets Directory", c.AssetsDirectory))

	return result
}

func (c *Config) ToJson() {
}
