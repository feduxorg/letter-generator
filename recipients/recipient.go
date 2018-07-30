package recipients

type Recipient struct {
	Name           string            `yaml:"name"`
	C_O            string            `yaml:"c_o"`
	Street         string            `yaml:"street"`
	City           string            `yaml:"city"`
	AdditionalInfo map[string]string `yaml:"additional_info"`
}
