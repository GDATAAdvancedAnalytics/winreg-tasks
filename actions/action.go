package actions

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

var (
	ErrUnknownPropertiesMagic = errors.New("unknown action properties magic")
)

type Actions struct {
	Context    string
	Properties []Properties
	Version    uint16
}

// FromBytes takes the content of an Actions Values and parses it.
func FromBytes(raw []byte) (*Actions, error) {
	stream := kaitai.NewStream(bytes.NewReader(raw))
	generatedActions := generated.NewActions()

	if err := generatedActions.Read(stream, nil, generatedActions); err != nil {
		return nil, err
	}

	if eof, err := stream.EOF(); !eof {
		return nil, fmt.Errorf("did not parse all data")
	} else if err != nil {
		return nil, fmt.Errorf("error trying to eof-check (%v)", err)
	}

	properties := make([]Properties, len(generatedActions.Actions))
	for i, action := range generatedActions.Actions {
		var props Properties
		var err error

		switch action.Magic {
		case uint16(ExecutionPropertiesMagic):
			props, err = NewExecutionProperties(action.Id.Str, action.Properties.(*generated.Actions_ExeTaskProperties))

		case uint16(ComHandlerPropertiesMagic):
			props, err = NewComHandlerProperties(action.Id.Str, action.Properties.(*generated.Actions_ComHandlerProperties))

		case uint16(EmailPropertiesMagic):
			props, err = NewEmailProperties(action.Id.Str, action.Properties.(*generated.Actions_EmailTaskProperties))

		case uint16(MessageboxPropertiesMagic):
			props, err = NewMessageboxProperties(action.Id.Str, action.Properties.(*generated.Actions_MessageboxTaskProperties))

		default:
			return nil, ErrUnknownPropertiesMagic
		}

		if err != nil {
			return nil, fmt.Errorf(`failed to parse action %d (%v)`, i, err)
		}

		properties[i] = props
	}

	actions := &Actions{
		Context:    generatedActions.Context.Str,
		Properties: properties,
		Version:    generatedActions.Version,
	}

	return actions, nil
}
