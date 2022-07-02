// SPDX-License-Identifier: MIT

package triggers

import (
	"fmt"
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
)

const SessionChangeTriggerMagic TriggerMagic = 0x7777

type SessionState int

const (
	ConsoleConnect    SessionState = 1
	ConsoleDisconnect SessionState = 2
	RemoteConnect     SessionState = 3
	RemoteDisconnect  SessionState = 4
	SessionLock       SessionState = 5
	SessionUnlock     SessionState = 6
)

type SessionChangeTrigger struct {
	GenericTriggerData *GenericTriggerData
	StateChange        SessionState
	User               *UserInfo
}

func NewSessionChangeTrigger(gen *generated.Triggers_SessionChangeTrigger, tz *time.Location) (*SessionChangeTrigger, error) {
	genericTriggerData, err := NewGenericTriggerData(gen.GenericData, tz)
	if err != nil {
		return nil, err
	}

	userInfo, err := NewUserInfo(gen.User)
	if err != nil {
		return nil, err
	}

	return &SessionChangeTrigger{
		GenericTriggerData: genericTriggerData,
		User:               userInfo,
		StateChange:        SessionState(gen.StateChange),
	}, nil
}

func IsSessionChangeTrigger(trigger Trigger) bool {
	return trigger.Magic() == SessionChangeTriggerMagic
}

func (t SessionChangeTrigger) Id() string {
	return t.GenericTriggerData.TriggerId
}

func (t SessionChangeTrigger) Magic() TriggerMagic {
	return SessionChangeTriggerMagic
}

func (t SessionChangeTrigger) Name() string {
	return "SessionChange"
}

func (t SessionChangeTrigger) String() string {
	return fmt.Sprintf(
		`<SessionChange user="%s" state="%d">`,
		t.User.UserToString(), t.StateChange,
	)
}
