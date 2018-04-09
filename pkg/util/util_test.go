package util

import (
	"testing"
	"time"
)

func TestPeriod(t *testing.T) {
	tcs := []struct {
		speed  int
		period time.Duration
		err    error
	}{
		{speed: 0, err: ErrSpeedInvalid},            // TC#1
		{speed: giga + 1, err: ErrspeedLimit},       // TC#2
		{speed: 10, period: time.Millisecond * 100}, // TC#3
	}

	for i, tc := range tcs {
		p, err := Period(tc.speed)
		if err != nil {
			if err != tc.err {
				t.Errorf("TC#%d: Period(%d), err %q, expected %q", i+1, tc.speed, err, tc.err)
			}
			continue
		}

		if p != tc.period {
			t.Errorf("TC#%d: Period(%d)=%s, expected %s", i+1, tc.speed, p, tc.period)
		}
	}
}
