// SPDX-License-Identifier: MIT

package triggers

import (
	"fmt"
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
)

const EventTriggerMagic TriggerMagic = 0xCCCC

type ValueQuery struct {
	Name  string
	Query string
}

type EventTrigger struct {
	GenericTriggerData *GenericTriggerData
	Subscription       string
	Unknown0           uint32
	Unknown1           uint32
	Unknown2           string
	ValueQueries       []ValueQuery
}

func NewEventTrigger(gen *generated.Triggers_EventTrigger, tz *time.Location) (*EventTrigger, error) {
	generic, err := NewGenericTriggerData(gen.GenericData, tz)
	if err != nil {
		return nil, err
	}

	valueQueries := make([]ValueQuery, gen.LenValueQueries.Value)
	for i, query := range gen.ValueQueries {
		valueQueries[i] = ValueQuery{
			Name:  query.Name.Content,
			Query: query.Value.Content,
		}
	}

	return &EventTrigger{
		GenericTriggerData: generic,
		Subscription:       gen.Subscription.Content,
		Unknown0:           gen.Unknown0,
		Unknown1:           gen.Unknown1,
		Unknown2:           gen.Unknown2.Content,
		ValueQueries:       valueQueries,
	}, nil
}

func IsEventTrigger(trigger Trigger) bool {
	return trigger.Magic() == EventTriggerMagic
}

func (t EventTrigger) Id() string {
	return t.GenericTriggerData.TriggerId
}

func (t EventTrigger) Magic() TriggerMagic {
	return EventTriggerMagic
}

func (t EventTrigger) Name() string {
	return "Event"
}

func (t EventTrigger) String() string {
	return fmt.Sprintf(`<Event subscription="%s">`, t.Subscription)
}
