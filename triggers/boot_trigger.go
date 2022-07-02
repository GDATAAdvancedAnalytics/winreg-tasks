// SPDX-License-Identifier: MIT

package triggers

import (
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
)

const BootTriggerMagic TriggerMagic = 0xFFFF

type BootTrigger struct {
	GenericData *GenericTriggerData
}

func NewBootTrigger(gen *generated.Triggers_BootTrigger, tz *time.Location) (*BootTrigger, error) {
	generic, err := NewGenericTriggerData(gen.GenericData, tz)
	if err != nil {
		return nil, err
	}

	return &BootTrigger{
		GenericData: generic,
	}, nil
}

func IsBootTrigger(trigger Trigger) bool {
	return trigger.Magic() == BootTriggerMagic
}

func (t BootTrigger) Id() string {
	return t.GenericData.TriggerId
}

func (t BootTrigger) Magic() TriggerMagic {
	return BootTriggerMagic
}

func (t BootTrigger) Name() string {
	return "Boot"
}

func (t BootTrigger) String() string {
	return `<Boot>`
}
