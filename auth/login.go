package auth

import (
	"log"

	"github.com/go-rod/rod"

	"linkedin-automation/stealth"
	"linkedin-automation/storage"
)

func Login(page *rod.Page, email string, password string) bool {
	log.Println("Starting login flow...")

	// --- Stealth: fingerprint masking (no-op for stability) ---
	stealth.MaskFingerprint(page)

	// --- Think time before interaction ---
	stealth.Think(500, 1200)

	// --- Email field (human typing) ---
	emailEl := page.MustElement("#username")
	stealth.HumanType(emailEl, email)

	// --- Think time between fields ---
	stealth.Think(400, 900)

	// --- Password field (human typing) ---
	passEl := page.MustElement("#password")
	stealth.HumanType(passEl, password)

	// --- Think time before clicking ---
	stealth.Think(600, 1000)

	// --- Login button (hover + click) ---
	loginBtn := page.MustElement("button[type='submit']")
	stealth.HumanClick(loginBtn)

	// --- Wait for navigation ---
	page.MustWaitLoad()

	// --- Stable success check ---
	_, err := page.Element("a[href='/logout']")
	if err == nil {
		log.Println("Login successful")
		_ = storage.SaveLogin("success")
		log.Println("Login state saved to state.json")
		return true

	}

	log.Println("Login failed")
	_ = storage.SaveLogin("failure") // ✅ STATE PERSISTENCE
	return false
}
