package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"go-core.course/homework-2/task-1/pkg/spider"
)

func main() {
	word := flag.String("word", "", "a word to search in page titles and urls")
	flag.Parse()
	if len(*word) == 0 {
		fmt.Println("empty word!")
		return
	}

	sites := []string{"https://yandex.ru", "https://ya.ru"}
	depth := 2

	// result is a map of sites mapped to url to page title
	hashMap := make(map[string]map[string]string)

	for _, site := range sites {
		data, err := spider.Scan(site, depth)
		if err != nil {
			log.Println(err)
			continue
		}
		hashMap[site] = data
	}

	// check for substring in urls and titles
	for source, v := range hashMap {
		for url, title := range v {

			if strings.Contains(url, *word) || strings.Contains(title, *word) {
				fmt.Printf("source url - %s. url - %s, title - %s\n", source, url, title)
			}
		}
	}
}
