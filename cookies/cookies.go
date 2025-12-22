package cookies

import (
	"encoding/json"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const cookieFile = "cookies.json"

func SaveCookies(page *rod.Page) error {
	cookies, err := page.Cookies([]string{})
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cookies, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(cookieFile, data, 0644)
}

func LoadCookies(page *rod.Page) error {
	data, err := os.ReadFile(cookieFile)
	if err != nil {
		return nil
	}

	var cookies []*proto.NetworkCookieParam
	if err := json.Unmarshal(data, &cookies); err != nil {
		return err
	}

	return page.SetCookies(cookies)
}
