package stealth

import (
	"math/rand"
	"time"
)

func Think(minMs int, maxMs int) {
	delay := minMs + rand.Intn(maxMs-minMs)
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
