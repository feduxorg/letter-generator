package converter

import (
	"github.com/fedux-org/letter-generator-go/recipients"
	"github.com/fedux-org/letter-generator-go/sender"
)

type TemplateContext struct {
	Recipient      *recipients.Recipient
	Sender         *sender.Sender
	Subject        string
	Signature      string
	Opening        string
	Closing        string
	HasAttachments bool
	HasPs          bool
}
