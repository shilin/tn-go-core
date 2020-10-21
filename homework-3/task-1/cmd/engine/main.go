package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	// we can import a substitution for a spider pkg, like:
	// "go-core.course/homework-3/task-1/pkg/dummySpider"
	"go-core.course/homework-3/task-1/pkg/spider"
)

type Scanner interface {
	Scan(string, int) (map[string]string, error)
}

type sauron struct {
}

// Here we can plug in a different Scan function from another package
func (s *sauron) Scan(url string, depth int) (map[string]string, error) {
	return spider.Scan(url, depth)
	// return dummySpider.Scan(url, depth)
}

func main() {
	sites := []string{"https://yandex.ru", "https://ya.ru"}
	depth := 2

	// result is a map of sites mapped to url to page title
	hashMap := make(map[string]map[string]string)

	var fetcher Scanner
	fetcher = &sauron{}

	fmt.Println("Please wait, scanning sites...")

	for _, site := range sites {
		data, err := fetcher.Scan(site, depth)
		if err != nil {
			log.Println(err)
			continue
		}
		hashMap[site] = data
	}

	fmt.Println(hashMap)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите слово для поиска>")
		word, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		word = strings.TrimSuffix(word, "\n")
		word = strings.TrimSuffix(word, "\r")
		if len(word) == 0 {
			fmt.Println("empty word!")
			continue
		}

		// check for substring in urls and titles
		for source, v := range hashMap {
			for url, title := range v {

				word = strings.ToLower(word)
				if strings.Contains(strings.ToLower(url), word) || strings.Contains(strings.ToLower(title), word) {
					fmt.Printf("source url - %s. url - %s, title - %s\n", source, url, title)
				}
			}
		}
	}
}
