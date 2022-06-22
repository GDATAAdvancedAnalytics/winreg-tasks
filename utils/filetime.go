package utils

import "time"

const secondsUntilEpoch = 11_644_473_600

func TimeFromFILETIME(filetime int64) time.Time {
	epoch := filetime/10_000_000 - secondsUntilEpoch
	return time.Unix(epoch, 0)
}
