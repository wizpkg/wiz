package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/wizpkg/scraper/api"
	"github.com/wizpkg/scraper/pypi"
)

var scrapers = map[string]api.ListPackagesFunction{}

func init() {
	pypiScraper := pypi.NewScraper()
	scrapers["pypi"] = pypiScraper.ListPackages
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var scrapeCmd = &cobra.Command{
	Use:   "scrape [registry/scraper name]",
	Short: "Scrape the specified package repository",

	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if fn, ok := scrapers[args[0]]; ok {
			fmt.Printf("Scraping from registry \"%s\"\n", args[0])
			packages := fn()

			file := fmt.Sprintf("packages-%s.txt", args[0])
			fmt.Printf("Done scraping. Writing to file: %s\n", file)

			f, err := os.Create(file)
			check(err)
			defer f.Close()

			w := bufio.NewWriter(f)
			defer w.Flush()

			i := 0
			for elem := range packages {
				_, err := w.WriteString(elem.Value.(string) + "\n")
				check(err)
				i++
			}
			fmt.Printf("Done! Total \"%s \"packages: %d\n", args[0], i)
		} else {
			fmt.Printf("Uh-oh, the scraper \"%s\" was not found", args[0])
		}
	},
}
