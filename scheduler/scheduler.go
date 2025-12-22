package scheduler

import "time"

func IsBusinessHours() bool {
	hour := time.Now().Hour()
	return hour >= 9 && hour <= 18
}
