package letter

import (
	"github.com/feduxorg/letter-generator-go/metadata"
	"github.com/feduxorg/letter-generator-go/recipients"
	"github.com/feduxorg/letter-generator-go/sender"
)

type Letter struct {
	Sender         sender.Sender
	Recipient      recipients.Recipient
	Subject        string
	Signature      string
	Opening        string
	Closing        string
	HasAttachments bool
	HasPs          bool
}

func New(sender sender.Sender, recipient recipients.Recipient, metadata metadata.Metadata) Letter {
	letter := Letter{}
	letter.Recipient = recipient
	letter.Sender = sender
	letter.Subject = metadata.Subject
	letter.Signature = metadata.Signature
	letter.Opening = metadata.Opening
	letter.Closing = metadata.Closing
	letter.HasAttachments = metadata.HasAttachments
	letter.HasPs = metadata.HasPs

	return letter
}
