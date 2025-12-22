package targeting

import (
	"github.com/go-rod/rod"
)

type Profile struct {
	Name     string
	Headline string
	URL      string
}

func ParseProfile(page *rod.Page, url string) Profile {
	page.MustNavigate(url)
	page.MustWaitLoad()

	name := ""
	headline := ""

	if el, err := page.Element("h1"); err == nil {
		name = el.MustText()
	}

	if el, err := page.Element("div.text-body-medium"); err == nil {
		headline = el.MustText()
	}

	return Profile{
		Name:     name,
		Headline: headline,
		URL:      url,
	}
}
