package recipients

type Recipient struct {
	Name           string            `json:"name"`
	C_O            string            `json:"c_o"`
	Street         string            `json:"street"`
	City           string            `json:"city"`
	AdditionalInfo map[string]string `json:"additional_info"`
}
