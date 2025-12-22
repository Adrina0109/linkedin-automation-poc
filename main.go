package main

import (
	"log"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	log.Println("Starting browser...")

	browser := rod.New().MustConnect()
	page := browser.MustPage("https://www.google.com")

	page.MustWaitLoad()

	log.Println("Typing into search box...")

	searchBox := page.MustElement("input[name='q']")
	searchBox.MustInput("LinkedIn Automation\n") // ← ENTER FIX

	log.Println("Search submitted")

	time.Sleep(4 * time.Second)

	browser.MustClose()
	log.Println("Browser closed")
}
