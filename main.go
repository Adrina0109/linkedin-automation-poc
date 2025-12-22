package main

import (
	"log"
	"time"

	"github.com/go-rod/rod"

	"linkedin-automation/auth"
	"linkedin-automation/config"
)

func main() {
	log.Println("Starting authentication demo...")

	cfg := config.Load()

	browser := rod.New().MustConnect()
	page := browser.MustPage("https://the-internet.herokuapp.com/login")
	page.MustWaitLoad()

	success := auth.Login(page, cfg.Email, cfg.Password)

	if !success {
		log.Println("Stopping due to login failure")
		browser.MustClose()
		return
	}

	log.Println("Login successful — keeping browser open for visibility")

	time.Sleep(10 * time.Second) // <-- this is the key

	browser.MustClose()
	log.Println("Browser closed cleanly")
}
