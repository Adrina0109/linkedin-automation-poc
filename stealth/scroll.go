package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func RandomScroll(page *rod.Page) {
	scrollY := rand.Intn(600) + 200
	page.MustEval(`window.scrollBy(0, arguments[0])`, scrollY)
	time.Sleep(time.Duration(500+rand.Intn(500)) * time.Millisecond)
}
