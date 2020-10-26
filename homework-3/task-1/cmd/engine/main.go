package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"go-core.course/homework-3/task-1/pkg/spider"
	// we can import a substitution for a spider pkg 
	// "go-core.course/homework-3/task-1/pkg/dummyFetch"
	"go-core.course/homework-3/task-1/pkg/scan"
)


func main() {
	sites := []string{"https://yandex.ru", "https://ya.ru"}
	depth := 2

	// s := new(dummyFetch.DummyFetch)
	s := new(spider.Spider)
	hashMap := scanResults(s, sites, depth)

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

func scanResults(s scan.Scanner, sites []string, depth int) map[string]map[string]string {
	// result is a map of sites mapped to url to page title
	hashMap := make(map[string]map[string]string)

	fmt.Println("Please wait, scanning sites...")

	for _, site := range sites {
		data, err := s.Scan(site, depth)
		if err != nil {
			log.Println(err)
			continue
		}
		hashMap[site] = data
	}

	fmt.Println(hashMap)
	return hashMap

}
