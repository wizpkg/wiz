package pypi

import (
	"fmt"

	utils "github.com/alexkreidler/goutils"
	"github.com/gocolly/colly"
)

// PyPi implements the Scraper interface
type PyPi struct{}

const simpleAPIURL = "https://pypi.org/simple/"
const jsonAPIURL = "https://pypi.org/pypi/"

// NewScraper returns a new scraper instance
func NewScraper() PyPi {
	return PyPi{}
}

// ListPackages returns a list of the registry's packages
func (p PyPi) ListPackages() <-chan utils.ConcurrentSliceItem {
	c := colly.NewCollector()

	packages := utils.NewConcurrentSlice()
	// Find and visit all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
		packages.Append(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(simpleAPIURL)
	c.Wait()
	return packages.Iter()
}
