package utils

import (
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
)

const secondsUntilEpoch = 11_644_473_600

func TimeFromFILETIME(filetime int64) time.Time {
	epoch := filetime/10_000_000 - secondsUntilEpoch
	return time.Unix(epoch, 0)
}

// TimeFromTSTime turns a generated TSTime object into a Golang time.Time.
// tz must be set to the Timezone of the TSTime object, otherwise the
// de-localization returns objects in the wrong timezone.
func TimeFromTSTime(gen *generated.Tstime, tz *time.Location) time.Time {
	tt := TimeFromFILETIME(int64(gen.Filetime.HighDateTime)<<32 | int64(gen.Filetime.LowDateTime))

	if gen.IsLocalized != 0 {
		// find out the offset to UTC by assuming tt was already in UTC and
		// translating it into the original timezone
		difference := tt.Sub(tt.In(tz))

		// now subtract the time difference so that we get the real UTC timestamp
		tt = tt.Add(-difference)
	}

	return tt
}

func DurationFromTSTimePeriod(gen *generated.Tstimeperiod) time.Duration {
	return time.Duration(gen.Year)*365*24*time.Hour + // TODO: check whether Microsoft really handles a year as 365 days internally
		time.Duration(gen.Month)*30*24*time.Hour + // TODO: check whether Microsoft really handles a month as 30 days internally
		time.Duration(gen.Day)*24*time.Hour +
		time.Duration(gen.Hour)*time.Hour +
		time.Duration(gen.Minute)*time.Minute +
		time.Duration(gen.Second)*time.Second
}
