package util

import (
	"fmt"
	"math"
	"time"
)

const (
	Separator = '#'

	giga = 1000 * 1000 * 1000
)

var (
	ErrSpeedInvalid = fmt.Errorf("should be > 0")
	ErrSpeedLimit   = fmt.Errorf("should be less than 10^9")
)

func Period(speed int) (time.Duration, error) {
	if speed <= 0 {
		return 0, ErrSpeedInvalid
	}
	if speed > giga {
		return 0, ErrSpeedLimit
	}

	return time.Nanosecond * time.Duration(math.Ceil(giga/float64(speed))), nil
}
