// SPDX-License-Identifier: MIT

package triggers

import (
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
)

const IdleTriggerMagic TriggerMagic = 0xEEEE

type IdleTrigger struct {
	GenericData *GenericTriggerData
}

func NewIdleTrigger(gen *generated.Triggers_IdleTrigger, tz *time.Location) (*IdleTrigger, error) {
	generic, err := NewGenericTriggerData(gen.GenericData, tz)
	if err != nil {
		return nil, err
	}

	return &IdleTrigger{
		GenericData: generic,
	}, nil
}

func IsIdleTrigger(trigger Trigger) bool {
	return trigger.Magic() == IdleTriggerMagic
}

func (t IdleTrigger) Id() string {
	return t.GenericData.TriggerId
}

func (t IdleTrigger) Magic() TriggerMagic {
	return IdleTriggerMagic
}

func (t IdleTrigger) Name() string {
	return "Idle"
}

func (t IdleTrigger) String() string {
	return `<Idle>`
}
