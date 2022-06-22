// SPDX-License-Identifier: MIT

package generated_test

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
)

func compareExeTask(task, expected *generated.Actions_Action) error {
	if task.Id.Str != expected.Id.Str {
		return errors.New("Task Id")
	}
	if task.Magic != expected.Magic {
		return errors.New("Task Magic")
	}
	t_props, ok := task.Properties.(*generated.Actions_ExeTaskProperties)
	if !ok {
		return errors.New("Task Type Cast")
	}

	e_props := expected.Properties.(*generated.Actions_ExeTaskProperties)

	if t_props.Arguments.Str != e_props.Arguments.Str {
		return errors.New("Task Props Arguments")
	}
	if t_props.Command.Str != e_props.Command.Str {
		return errors.New("Task Props Command")
	}
	if t_props.WorkingDirectory.Str != e_props.WorkingDirectory.Str {
		return errors.New("Task Props WorkingDirectory")
	}
	if t_props.Flags != e_props.Flags {
		return errors.New("Task Props HideAppWindow")
	}
	return nil
}

func compareComHandlerTask(task, expected *generated.Actions_Action) error {
	if task.Id.Str != expected.Id.Str {
		return errors.New("Task Id")
	}
	if task.Magic != expected.Magic {
		return errors.New("Task Magic")
	}
	t_props, ok := task.Properties.(*generated.Actions_ComHandlerProperties)
	if !ok {
		return errors.New("Task Type Cast")
	}

	e_props := expected.Properties.(*generated.Actions_ComHandlerProperties)

	if !bytes.Equal(t_props.Clsid, e_props.Clsid) {
		return errors.New("Task Props Arguments")
	}
	if t_props.Data.Str != e_props.Data.Str {
		return errors.New("Task Props Command")
	}
	return nil
}

func compareMessageBoxTask(task, expected *generated.Actions_Action) error {
	if task.Id.Str != expected.Id.Str {
		return errors.New("Task Id")
	}
	if task.Magic != expected.Magic {
		return errors.New("Task Magic")
	}
	t_props, ok := task.Properties.(*generated.Actions_MessageboxTaskProperties)
	if !ok {
		return errors.New("Task Type Cast")
	}

	e_props := expected.Properties.(*generated.Actions_MessageboxTaskProperties)

	if t_props.Caption.Str != e_props.Caption.Str {
		return errors.New("Task Props Caption")
	}
	if t_props.Content.Str != e_props.Content.Str {
		return errors.New("Task Props Content")
	}
	return nil
}

func compareEmailTask(task, expected *generated.Actions_Action) error {
	if task.Id.Str != expected.Id.Str {
		return errors.New("Task Id")
	}
	if task.Magic != expected.Magic {
		return errors.New("Task Magic")
	}

	props, ok := task.Properties.(*generated.Actions_EmailTaskProperties)
	if !ok {
		return errors.New("Task Type conversion error")
	}

	e_props := expected.Properties.(*generated.Actions_EmailTaskProperties)

	if props.Server.Str != e_props.Server.Str {
		return errors.New("Server")
	}
	if props.Subject.Str != e_props.Subject.Str {
		return errors.New("Subject")
	}
	if props.To.Str != e_props.To.Str {
		return errors.New("To")
	}
	if props.Cc.Str != e_props.Cc.Str {
		return errors.New("Cc")
	}
	if props.Bcc.Str != e_props.Bcc.Str {
		return errors.New("Bcc")
	}
	if props.ReplyTo.Str != e_props.ReplyTo.Str {
		return errors.New("ReplyTo")
	}
	if props.From.Str != e_props.From.Str {
		return errors.New("From")
	}
	if props.Body.Str != e_props.Body.Str {
		return errors.New("Body")
	}
	if len(props.Headers) != len(e_props.Headers) {
		return errors.New("Headers Length")
	}
	for i := range props.Headers {
		if props.Headers[i].Key.Str != e_props.Headers[i].Key.Str {
			return fmt.Errorf("Header %d Name mismatch", i)
		}
		if props.Headers[i].Value.Str != e_props.Headers[i].Value.Str {
			return fmt.Errorf("Header %d Value mismatch", i)
		}
	}
	if len(props.AttachmentFilenames) != len(e_props.AttachmentFilenames) {
		return errors.New("Attachments Length")
	}
	for i := range props.AttachmentFilenames {
		if props.AttachmentFilenames[i].Str != e_props.AttachmentFilenames[i].Str {
			return fmt.Errorf("Attachment %d mismatch", i)
		}
	}
	return nil
}
