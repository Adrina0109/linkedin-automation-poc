package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func HumanType(el *rod.Element, text string) {
	for _, ch := range text {
		el.MustInput(string(ch))

		// Variable typing speed (human-like)
		delay := time.Duration(80+rand.Intn(120)) * time.Millisecond
		time.Sleep(delay)
	}
}
