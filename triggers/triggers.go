// SPDX-License-Identifier: MIT

package triggers

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

var (
	ErrUnknownTriggerMagic = errors.New("unknown trigger magic")
)

type TriggerMagic uint16

type Trigger interface {
	// The ID of the Trigger. Usually empty.
	Id() string
	// The Magic of the Trigger
	Magic() TriggerMagic
	// The short name of the Trigger
	Name() string
	// A String representation of the Trigger including its most important values
	String() string
}

type Triggers struct {
	Header    *Header
	JobBucket *JobBucket
	Triggers  []Trigger
}

// FromBytes takes the content of an Triggers Values and parses it.
func FromBytes(raw []byte, tz *time.Location) (*Triggers, error) {
	stream := kaitai.NewStream(bytes.NewReader(raw))
	gen := generated.NewTriggers()

	if err := gen.Read(stream, nil, gen); err != nil {
		return nil, err
	}

	if eof, err := stream.EOF(); !eof {
		return nil, fmt.Errorf("did not parse all data")
	} else if err != nil {
		return nil, fmt.Errorf("error trying to eof-check (%v)", err)
	}

	header, err := NewHeader(gen.Header, tz)
	if err != nil {
		return nil, err
	}

	bucket, err := NewJobBucket(gen.JobBucket)
	if err != nil {
		return nil, err
	}

	triggers := &Triggers{
		Header:    header,
		JobBucket: bucket,
	}

	for _, trig := range gen.Triggers {
		var t Trigger
		switch trig.Magic.Value {
		case uint32(BootTriggerMagic):
			t, err = NewBootTrigger(trig.Properties.(*generated.Triggers_BootTrigger), tz)
		case uint32(EventTriggerMagic):
			t, err = NewEventTrigger(trig.Properties.(*generated.Triggers_EventTrigger), tz)
		case uint32(IdleTriggerMagic):
			t, err = NewIdleTrigger(trig.Properties.(*generated.Triggers_IdleTrigger), tz)
		case uint32(LogonTriggerMagic):
			t, err = NewLogonTrigger(trig.Properties.(*generated.Triggers_LogonTrigger), tz)
		case uint32(RegistrationTriggerMagic):
			t, err = NewRegistrationTrigger(trig.Properties.(*generated.Triggers_RegistrationTrigger), tz)
		case uint32(SessionChangeTriggerMagic):
			t, err = NewSessionChangeTrigger(trig.Properties.(*generated.Triggers_SessionChangeTrigger), tz)
		case uint32(TimeTriggerMagic):
			t, err = NewTimeTrigger(trig.Properties.(*generated.Triggers_TimeTrigger), tz)
		case uint32(WnfStateChangeTriggerMagic):
			t, err = NewWnfStateChangeTrigger(trig.Properties.(*generated.Triggers_WnfStateChangeTrigger), tz)
		default:
			return nil, ErrUnknownTriggerMagic
		}

		if err != nil {
			return nil, err
		}
		triggers.Triggers = append(triggers.Triggers, t)
	}

	return triggers, nil
}
