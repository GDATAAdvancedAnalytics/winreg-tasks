package actions

import (
	"fmt"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
)

const EmailPropertiesMagic PropertiesMagic = 0x8888

type EmailHeader struct {
	Name  string
	Value string
}

type EmailProperties struct {
	Id string

	From                string
	To                  string
	Cc                  string
	Bcc                 string
	ReplyTo             string
	Server              string
	Subject             string
	Body                string
	AttachmentFilenames []string
	Headers             []EmailHeader
}

func NewEmailProperties(id string, gen *generated.Actions_EmailTaskProperties) (*EmailProperties, error) {
	attachmentFilenames := make([]string, gen.NumAttachmentFilenames)
	for i, file := range gen.AttachmentFilenames {
		attachmentFilenames[i] = file.Str
	}

	headers := make([]EmailHeader, gen.NumHeaders)
	for i, header := range gen.Headers {
		headers[i] = EmailHeader{Name: header.Key.Str, Value: header.Value.Str}
	}

	return &EmailProperties{
		Id:                  id,
		From:                gen.From.Str,
		To:                  gen.To.Str,
		Cc:                  gen.Cc.Str,
		Bcc:                 gen.Bcc.Str,
		ReplyTo:             gen.ReplyTo.Str,
		Server:              gen.Server.Str,
		Subject:             gen.Subject.Str,
		Body:                gen.Body.Str,
		AttachmentFilenames: attachmentFilenames,
		Headers:             headers,
	}, nil
}

func IsEmailProperties(properties Properties) bool {
	return properties.Magic() == EmailPropertiesMagic
}

func (e EmailProperties) Magic() PropertiesMagic {
	return EmailPropertiesMagic
}

func (e EmailProperties) Name() string {
	return "Email"
}

func (e EmailProperties) String() string {
	return fmt.Sprintf(
		`<Email id="%s" from="%s" to="%s" subject="%s" ...>`,
		e.Id, e.From, e.To, e.Subject,
	)
}
