package main

import (
	"log"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	log.Println("Starting browser...")

	browser := rod.New().MustConnect()
	page := browser.MustPage("https://duckduckgo.com")

	page.MustWaitLoad()

	log.Println("Typing into search box...")

	// DuckDuckGo search input
	searchBox := page.MustElement("#searchbox_input")

	// Type query + Enter
	searchBox.MustInput("LinkedIn Automation\n")

	log.Println("Search submitted")

	// Wait to observe results
	time.Sleep(4 * time.Second)

	browser.MustClose()
	log.Println("Browser closed")
}
