package actions

import (
	"fmt"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
	"github.com/google/uuid"
)

type PropertiesMagic uint16

const (
	ExecutionPropertiesMagic  PropertiesMagic = 0x6666
	ComHandlerPropertiesMagic PropertiesMagic = 0x7777
	EmailPropertiesMagic      PropertiesMagic = 0x8888
	MessageboxPropertiesMagic PropertiesMagic = 0x9999
)

type Properties interface {
	Magic() PropertiesMagic
	Name() string
	String() string
}

type ExecutionProperties struct {
	Id string

	Arguments        string
	Command          string
	WorkingDirectory string

	Flags uint16
}

func NewExecutionProperties(id string, gen *generated.Actions_ExeTaskProperties) (*ExecutionProperties, error) {
	return &ExecutionProperties{
		Id:               id,
		Arguments:        gen.Arguments.Str,
		Command:          gen.Command.Str,
		WorkingDirectory: gen.WorkingDirectory.Str,
		Flags:            gen.Flags,
	}, nil
}

func (e ExecutionProperties) Magic() PropertiesMagic {
	return ExecutionPropertiesMagic
}

func (e ExecutionProperties) Name() string {
	return "Execution"
}

func (e ExecutionProperties) String() string {
	return fmt.Sprintf(
		`<Execution id="%s" command="%s" arguments="%s" workingDirectory="%s" flags=0x{%02x}`,
		e.Id, e.Command, e.Arguments, e.WorkingDirectory, e.Flags,
	)
}

type ComHandlerProperties struct {
	Id string

	Clsid uuid.UUID
	Data  string
}

func NewComHandlerProperties(id string, gen *generated.Actions_ComHandlerProperties) (*ComHandlerProperties, error) {
	clsid, err := utils.UuidFromMemory(gen.Clsid)
	if err != nil {
		return nil, err
	}

	return &ComHandlerProperties{
		Id:    id,
		Clsid: clsid,
		Data:  gen.Data.Str,
	}, nil
}

func (c ComHandlerProperties) Magic() PropertiesMagic {
	return ComHandlerPropertiesMagic
}

func (e ComHandlerProperties) Name() string {
	return "ComHandler"
}

func (c ComHandlerProperties) String() string {
	return fmt.Sprintf(
		`<ComHandler id="%s" clsid="%s" data="%s">`,
		c.Id, c.Clsid.String(), c.Data,
	)
}

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

type MessageboxProperties struct {
	Id string

	Caption string
	Content string
}

func NewMessageboxProperties(id string, gen *generated.Actions_MessageboxTaskProperties) (*MessageboxProperties, error) {
	return &MessageboxProperties{
		Id:      id,
		Caption: gen.Caption.Str,
		Content: gen.Content.Str,
	}, nil
}

func (m MessageboxProperties) Magic() PropertiesMagic {
	return MessageboxPropertiesMagic
}

func (e MessageboxProperties) Name() string {
	return "Messagebox"
}

func (m MessageboxProperties) String() string {
	return fmt.Sprintf(
		`<Messagebox id="%s" caption="%s" content="%s">`,
		m.Id, m.Caption, m.Content,
	)
}
