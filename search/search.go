package search

import (
	"log"
	"github.com/go-rod/rod"
)

func SearchPeople(page *rod.Page, query string, maxScrolls int) []string {
	log.Println("Starting LinkedIn people search...")

	searchURL := "https://www.linkedin.com/search/results/people/?keywords=" + query
	page.MustNavigate(searchURL)
	page.MustWaitLoad()

	profiles := make(map[string]bool)

	for i := 0; i < maxScrolls; i++ {
		page.MustEval(`window.scrollBy(0, document.body.scrollHeight)`)
		page.MustWaitIdle()

		links := page.MustElements("a[href*='/in/']")
		for _, l := range links {
			href := l.MustProperty("href").String()
			profiles[href] = true
		}
	}

	var result []string
	for p := range profiles {
		result = append(result, p)
	}

	log.Printf("Collected %d unique profiles\n", len(result))
	return result
}
