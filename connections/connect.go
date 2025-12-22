package connections

import (
	"log"
	"time"
	"github.com/go-rod/rod"
)

var DailyLimit = 5

func SendConnection(page *rod.Page, sent *int) {
	if *sent >= DailyLimit {
		log.Println("Daily connection limit reached")
		return
	}

	btn, err := page.Element("button:contains('Connect')")
	if err != nil {
		log.Println("Connect button not found")
		return
	}

	btn.MustHover()
	time.Sleep(800 * time.Millisecond)
	btn.MustClick()

	*sent++
	time.Sleep(3 * time.Second)
}
