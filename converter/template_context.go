package converter

import (
	"github.com/feduxorg/letter-generator/recipients"
	"github.com/feduxorg/letter-generator/sender"
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
