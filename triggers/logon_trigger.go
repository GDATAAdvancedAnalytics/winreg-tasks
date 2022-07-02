// SPDX-License-Identifier: MIT

package triggers

import (
	"fmt"
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
)

const LogonTriggerMagic TriggerMagic = 0xAAAA

type LogonTrigger struct {
	GenericData *GenericTriggerData
	User        *UserInfo
}

func NewLogonTrigger(gen *generated.Triggers_LogonTrigger, tz *time.Location) (*LogonTrigger, error) {
	generic, err := NewGenericTriggerData(gen.GenericData, tz)
	if err != nil {
		return nil, err
	}

	userInfo, err := NewUserInfo(gen.User)
	if err != nil {
		return nil, err
	}

	return &LogonTrigger{
		GenericData: generic,
		User:        userInfo,
	}, nil
}

func IsLogonTrigger(trigger Trigger) bool {
	return trigger.Magic() == LogonTriggerMagic
}

func (t LogonTrigger) Id() string {
	return t.GenericData.TriggerId
}

func (t LogonTrigger) Magic() TriggerMagic {
	return LogonTriggerMagic
}

func (t LogonTrigger) Name() string {
	return "Logon"
}

func (t LogonTrigger) String() string {
	return fmt.Sprintf(
		`<Logon user="%s">`,
		t.User.UserToString(),
	)
}
