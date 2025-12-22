package main

import (
	"log"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	log.Println("Starting browser...")

	browser := rod.New().MustConnect()
	page := browser.MustPage("https://example.com")
	page.MustWaitLoad()

	log.Println("Hovering over heading...")

	// Find the heading
	heading := page.MustElement("h1")

	// Hover (human-like behavior)
	heading.MustHover()
	time.Sleep(1 * time.Second)

	log.Println("Reading page text...")

	// Read text content
	text := heading.MustText()
	log.Println("Heading text:", text)

	time.Sleep(2 * time.Second)

	browser.MustClose()
	log.Println("Browser closed")
}