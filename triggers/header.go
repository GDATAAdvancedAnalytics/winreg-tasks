// SPDX-License-Identifier: MIT

package triggers

import (
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
)

type Header struct {
	// The version of the structure
	Version uint8
	// The earliest StartBoundary of all triggers
	StartBoundary time.Time
	// The latest EndBoundary of all triggers
	EndBoundary time.Time
}

// NewHeader converts a generated Triggers_Headers into a more accessible Header object
//
// The tz must be set to the timezone where the the Trigger data were originally created in.
// Otherwise there might be differences in the times represented by this code and what Windows
// shows in the Task Scheduler MMC snap-in.
func NewHeader(gen *generated.Triggers_Header, tz *time.Location) (*Header, error) {
	return &Header{
		Version:       gen.Version.Value,
		StartBoundary: utils.TimeFromTSTime(gen.StartBoundary, tz),
		EndBoundary:   utils.TimeFromTSTime(gen.EndBoundary, tz),
	}, nil
}
