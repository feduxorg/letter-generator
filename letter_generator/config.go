package letter_generator

import (
	"fmt"
	"strings"
)

type Config struct {
	RemoteSources  []string `json:remote_sources`
	RecipientsFile string   `json:recipients_file`
	MetadataFile   string   `json:metadata_file`
	SenderFile     string   `json:sender_file`
	TemplateFile   string   `json:template_file`
}

func (c *Config) ToString() []string {
	result := []string{}
	result = append(result, fmt.Sprintf("%20s | %-30s", "Option", "Value"))
	result = append(result, fmt.Sprintf("%20s | %-30s", "Remote Sources", strings.Join(c.RemoteSources, ", ")))
	result = append(result, fmt.Sprintf("%20s | %-30s", "Recipients File", c.RecipientsFile))
	result = append(result, fmt.Sprintf("%20s | %-30s", "Metadata File", c.MetadataFile))
	result = append(result, fmt.Sprintf("%20s | %-30s", "Sender File", c.SenderFile))
	result = append(result, fmt.Sprintf("%20s | %-30s", "Template File", c.TemplateFile))

	return result
}

func (c *Config) ToJson() {
}
