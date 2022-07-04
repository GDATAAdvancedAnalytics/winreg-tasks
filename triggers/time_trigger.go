// SPDX-License-Identifier: MIT

package triggers

import (
	"fmt"
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
)

const TimeTriggerMagic TriggerMagic = 0xDDDD

type TimeTrigger struct {
	JobSchedule *JobSchedule

	id string
}

func NewTimeTrigger(gen *generated.Triggers_TimeTrigger, tz *time.Location) (*TimeTrigger, error) {
	schedule, err := NewJobSchedule(gen.JobSchedule, tz)
	if err != nil {
		return nil, err
	}

	triggerId := ""
	if gen.TriggerId != nil {
		triggerId = gen.TriggerId.Str
	}

	return &TimeTrigger{
		JobSchedule: schedule,
		id:          triggerId,
	}, nil
}

func IsTimeTrigger(trigger Trigger) bool {
	return trigger.Magic() == TimeTriggerMagic
}

func (t TimeTrigger) Id() string {
	return t.id
}

func (t TimeTrigger) Magic() TriggerMagic {
	return TimeTriggerMagic
}

func (t TimeTrigger) Name() string {
	return "Time"
}

func (t TimeTrigger) String() string {
	switch t.JobSchedule.Mode {
	case OneTime:
		return fmt.Sprintf(
			`<Time start="%s">`,
			t.JobSchedule.StartBoundary,
		)
	case Daily:
		return fmt.Sprintf(
			`<Time start="%s" mode=%s repeat_every_n_days=%d>`,
			t.JobSchedule.StartBoundary.String(), TimeModeToString(t.JobSchedule.Mode),
			t.JobSchedule.RepeatEvery,
		)
	case Weekly:
		return fmt.Sprintf(
			`<Time start="%s" mode=%s days_of_week="%s" repeat_every_n_weeks=%d>`,
			t.JobSchedule.StartBoundary.String(), TimeModeToString(t.JobSchedule.Mode),
			utils.BitmapToString(uint64(t.JobSchedule.DaysOfWeek)), t.JobSchedule.RepeatEvery,
		)
	case DaysInMonths:
		return fmt.Sprintf(
			`<Time start="%s" mode=%s months="%s" days_on_month="%s">`,
			t.JobSchedule.StartBoundary.String(), TimeModeToString(t.JobSchedule.Mode),
			utils.BitmapToString(uint64(t.JobSchedule.Months)),
			utils.BitmapToString(uint64(t.JobSchedule.DaysInMonth)),
		)
	case DaysInWeeksInMonths:
		return fmt.Sprintf(
			`<Time start="%s" mode=%s months="%s" weeks_in_month="%s" days_of_week="%s">`,
			t.JobSchedule.StartBoundary.String(), TimeModeToString(t.JobSchedule.Mode),
			utils.BitmapToString(uint64(t.JobSchedule.Months)),
			utils.BitmapToString(uint64(t.JobSchedule.WeeksInMonth)),
			utils.BitmapToString(uint64(t.JobSchedule.DaysOfWeek)),
		)
	default:
		return fmt.Sprintf(
			`<Time start="%s" mode=%s>`,
			t.JobSchedule.StartBoundary.String(), TimeModeToString(t.JobSchedule.Mode),
		)
	}
}
