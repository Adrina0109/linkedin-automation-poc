package main

import (
	"log"
	"strings"
	"time"

	"github.com/go-rod/rod"

	"linkedin-automation/auth"
	"linkedin-automation/config"
	"linkedin-automation/cookies"
	"linkedin-automation/connections"
	"linkedin-automation/storage"
)

func main() {
	log.Println("Starting LinkedIn automation demo...")

	cfg := config.Load()

	browser := rod.New().MustConnect()
	browser = browser.MustIncognito()

	page := browser.MustPage("https://the-internet.herokuapp.com/login")
	page.MustWaitLoad()

	success := auth.Login(page, cfg.Email, cfg.Password)
	if !success {
		log.Println("Stopping due to login failure")
		browser.MustClose()
		return
	}

	if els, _ := page.Elements("iframe"); len(els) > 0 {
		for _, el := range els {
			src, _ := el.Attribute("src")
			if src != nil && *src != "" {
				log.Println("Possible captcha iframe detected. Stopping automation.")
				browser.MustClose()
				return
			}
		}
	}

	bodyText := page.MustElement("body").MustText()

	if strings.Contains(bodyText, "checkpoint") {
		log.Println("Security checkpoint detected. Stopping automation.")
		browser.MustClose()
		return
	}

	if strings.Contains(bodyText, "unusual activity") {
		log.Println("Unusual activity warning detected. Stopping automation.")
		browser.MustClose()
		return
	}

	log.Println("Saving session cookies...")
	if err := cookies.SaveCookies(page); err != nil {
		log.Println("Failed to save cookies:", err)
	} else {
		log.Println("Cookies saved to cookies.json")
	}

	sent := 0
	log.Println("Attempting connection request (demo)...")
	connections.SendConnection(page, &sent, cfg)

	storage.SaveReport(storage.Report{
		Timestamp:       time.Now(),
		Logins:          1,
		ConnectionsSent: sent,
	})

	log.Println("Automation completed — keeping browser open for visibility")
	time.Sleep(10 * time.Second)

	browser.MustClose()
	log.Println("Browser closed cleanly")
}
