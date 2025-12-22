package stealth

import (
	"time"

	"github.com/go-rod/rod"
)

func HumanClick(el *rod.Element) {
	el.MustHover()
	time.Sleep(300 * time.Millisecond)
	el.MustClick()
}
