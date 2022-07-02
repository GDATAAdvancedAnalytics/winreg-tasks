// SPDX-License-Identifier: MIT

package triggers

import (
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
)

type TimeMode int

const (
	// runs at JobSchedule.Start
	OneTime TimeMode = 0
	// runs at JobSchedule.Start and repeats every JobSchedule.RepeatEvery days
	Daily TimeMode = 1
	// run on JobSchedule.DaysOfWeek every JobSchedule.RepeatEvery weeks starting at JobSchedule.Start
	Weekly TimeMode = 2
	// run in JobSchedule.Months on JobSchedule.DaysInMonth starting at JobSchedule.Start
	DaysInMonths TimeMode = 3
	// run in JobSchedule.Months in JobSchedule.WeeksInMonth on JobSchedule.DaysOfWeek starting at JobSchedule.Start
	DaysInWeeksInMonths TimeMode = 4
)

type DayOfWeek int

const (
	Sunday    DayOfWeek = 0x01
	Monday    DayOfWeek = 0x02
	Tuesday   DayOfWeek = 0x04
	Wednesday DayOfWeek = 0x08
	Thursday  DayOfWeek = 0x10
	Friday    DayOfWeek = 0x20
	Saturday  DayOfWeek = 0x40
)

type Months int

const (
	January   Months = 0x001
	February  Months = 0x002
	March     Months = 0x004
	April     Months = 0x008
	May       Months = 0x010
	June      Months = 0x020
	July      Months = 0x040
	August    Months = 0x080
	September Months = 0x100
	October   Months = 0x200
	November  Months = 0x400
	December  Months = 0x800
)

type JobSchedule struct {
	StartBoundary      time.Time
	EndBoundary        time.Time
	Unknown0           time.Time
	RepetitionInterval time.Duration
	RepetitionDuration time.Duration
	ExecutionTimeLimit time.Duration
	Mode               TimeMode
	// value depends on Mode setting
	DaysOfWeek DayOfWeek
	// value depends on Mode setting
	Months Months
	// bitmap of days in month; first day is 0x1, second is 0x2, third is 0x4, and so on
	DaysInMonth uint32
	// bitmap of weeks in month; first week is 0x1, second is 0x2, third is 0x4, and so on
	WeeksInMonth uint16
	// meaning of value depends on mode; check TimeMode constants
	RepeatEvery            uint
	StopTasksAtDurationEnd bool
	Enabled                bool
	Unknown1               uint32
	MaxDelay               time.Duration
}

func NewJobSchedule(gen *generated.JobSchedule, tz *time.Location) (*JobSchedule, error) {
	schedule := &JobSchedule{
		StartBoundary:          utils.TimeFromTSTime(gen.StartBoundary, tz),
		EndBoundary:            utils.TimeFromTSTime(gen.EndBoundary, tz),
		Unknown0:               utils.TimeFromTSTime(gen.Unknown0, tz),
		RepetitionInterval:     time.Second * time.Duration(gen.RepetitionIntervalSeconds),
		RepetitionDuration:     time.Second * time.Duration(gen.RepetitionDurationSeconds),
		ExecutionTimeLimit:     time.Second * time.Duration(gen.ExecutionTimeLimitSeconds),
		Mode:                   TimeMode(gen.Mode),
		StopTasksAtDurationEnd: gen.StopTasksAtDurationEnd != 0,
		Enabled:                gen.IsEnabled != 0,
		Unknown1:               gen.Unknown1,
		MaxDelay:               time.Second * time.Duration(gen.MaxDelaySeconds),
	}

	switch schedule.Mode {
	case OneTime:
		break
	case Daily:
		schedule.RepeatEvery = uint(gen.Data1)
	case Weekly:
		schedule.RepeatEvery = uint(gen.Data1)
		schedule.DaysOfWeek = DayOfWeek(gen.Data2)
	case DaysInMonths:
		schedule.Months = Months(gen.Data3)
		schedule.DaysInMonth = uint32(gen.Data2)<<16 | uint32(gen.Data1)
	case DaysInWeeksInMonths:
		schedule.DaysOfWeek = DayOfWeek(gen.Data1)
		schedule.WeeksInMonth = gen.Data2
		schedule.Months = Months(gen.Data3)
	}

	return schedule, nil
}
