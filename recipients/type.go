package recipients

type Recipient struct {
	Name           string   `json:"name"`
	C_O            string   `json:"c_o"`
	Street         string   `json:"street"`
	City           string   `json:"city"`
	AdditionalInfo []string `json:"additional_info"`
}

type RecipientManager struct {
	Recipients []Recipient
}
