package stealth

import (
	"time"
)

var lastAction time.Time

func RateLimit(minGapSeconds int) {
	if time.Since(lastAction) < time.Duration(minGapSeconds)*time.Second {
		time.Sleep(time.Duration(minGapSeconds)*time.Second)
	}
	lastAction = time.Now()
}
