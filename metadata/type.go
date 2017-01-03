package metadata

type Metadata struct {
	Subject        string `json:"subject"`
	Signature      string `json:"signature"`
	Opening        string `json:"opening"`
	Closing        string `json:"closing"`
	HasAttachments bool   `json:"has_attachments"`
	HasPs          bool   `json:"has_ps"`
}
