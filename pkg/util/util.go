package util

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

const (
	Separator = '#'

	giga = 1000 * 1000 * 1000
)

var (
	NumberLimit big.Int
)

func init() {
	NumberLimit.Exp(big.NewInt(256), big.NewInt(16), nil)
}

func Period(speed int) (time.Duration, error) {
	if speed <= 0 {
		return 0, fmt.Errorf("should be > 0")
	}
	if speed > giga {
		return 0, fmt.Errorf("should be less than 10^9")
	}

	return time.Nanosecond * time.Duration(math.Ceil(giga/float64(speed))), nil
}
