// SPDX-License-Identifier: MIT

package triggers

import (
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
)

type GenericTriggerData struct {
	StartBoundary       time.Time
	EndBoundary         time.Time
	Delay               time.Duration
	Timeout             time.Duration
	RepetitionInterval  time.Duration
	RepetitionDuration  time.Duration
	RepetitionDuration2 time.Duration
	StopAtDurationEnd   bool
	Enabled             bool
	Unknown             []byte
	TriggerId           string
}

func NewGenericTriggerData(gen *generated.Triggers_GenericTriggerData, tz *time.Location) (*GenericTriggerData, error) {
	triggerId := ""
	if gen.TriggerId != nil {
		triggerId = gen.TriggerId.Str
	}

	return &GenericTriggerData{
		StartBoundary:       utils.TimeFromTSTime(gen.StartBoundary, tz),
		EndBoundary:         utils.TimeFromTSTime(gen.EndBoundary, tz),
		Delay:               time.Duration(gen.DelaySeconds) * time.Second,
		Timeout:             time.Duration(gen.TimeoutSeconds) * time.Second,
		RepetitionInterval:  time.Duration(gen.RepetitionIntervalSeconds) * time.Second,
		RepetitionDuration:  time.Duration(gen.RepetitionDurationSeconds) * time.Second,
		RepetitionDuration2: time.Duration(gen.RepetitionDurationSeconds2) * time.Second,
		StopAtDurationEnd:   gen.StopAtDurationEnd != 0,
		Enabled:             gen.Enabled.Value != 0,
		Unknown:             gen.Unknown[:],
		TriggerId:           triggerId,
	}, nil
}
