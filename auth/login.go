package auth

import (
	"log"

	"github.com/go-rod/rod"

	"linkedin-automation/stealth"
	"linkedin-automation/storage"
)

func Login(page *rod.Page, email string, password string) bool {
	log.Println("Starting login flow...")

	stealth.MaskFingerprint(page)

	stealth.Think(500, 1200)

	emailEl := page.MustElement("#username")
	stealth.HumanType(emailEl, email)

	stealth.Think(400, 900)

	passEl := page.MustElement("#password")
	stealth.HumanType(passEl, password)

	stealth.Think(600, 1000)

	loginBtn := page.MustElement("button[type='submit']")
	stealth.HumanClick(loginBtn)

	page.MustWaitLoad()

	_, err := page.Element("a[href='/logout']")
	if err == nil {
		log.Println("Login successful")
		_ = storage.SaveLogin("success")
		log.Println("Login state saved to state.json")
		return true
	}

	log.Println("Login failed")
	_ = storage.SaveLogin("failure")
	return false
}
