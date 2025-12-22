package stealth

import "github.com/go-rod/rod"

// MaskFingerprint is intentionally a no-op for stability.
// Navigator modifications are avoided to prevent Rod eval crashes.
func MaskFingerprint(page *rod.Page) {
	// intentionally left blank
}
