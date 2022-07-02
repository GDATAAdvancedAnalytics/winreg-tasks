// SPDX-License-Identifier: MIT

package triggers

import (
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
)

const RegistrationTriggerMagic TriggerMagic = 0x8888

type RegistrationTrigger struct {
	GenericData *GenericTriggerData
}

func NewRegistrationTrigger(gen *generated.Triggers_RegistrationTrigger, tz *time.Location) (*RegistrationTrigger, error) {
	generic, err := NewGenericTriggerData(gen.GenericData, tz)
	if err != nil {
		return nil, err
	}

	return &RegistrationTrigger{
		GenericData: generic,
	}, nil
}

func IsRegistrationTrigger(trigger Trigger) bool {
	return trigger.Magic() == RegistrationTriggerMagic
}

func (t RegistrationTrigger) Id() string {
	return t.GenericData.TriggerId
}

func (t RegistrationTrigger) Magic() TriggerMagic {
	return RegistrationTriggerMagic
}

func (t RegistrationTrigger) Name() string {
	return "Registration"
}

func (t RegistrationTrigger) String() string {
	return `<Registration>`
}
