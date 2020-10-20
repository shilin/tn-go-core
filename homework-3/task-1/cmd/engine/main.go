package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"go-core.course/homework-2/task-1/pkg/spider"
)

type scanner interface {
	Scan(string, int) (map[string]string, error)
}

type sauron struct {
}

func (s *sauron) Scan(url string, depth int) (map[string]string, error) {
	return spider.Scan(url, depth)
}

func main() {
	sites := []string{"https://yandex.ru", "https://ya.ru"}
	depth := 2

	// result is a map of sites mapped to url to page title
	hashMap := make(map[string]map[string]string)

	fmt.Println("Please wait, scanning sites...")

	var fetcher scanner
	fetcher = &sauron{}

	for _, site := range sites {
		data, err := fetcher.Scan(site, depth)
		if err != nil {
			log.Println(err)
		}
		hashMap[site] = data
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите слово для поиска>")
		word, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		word = strings.TrimSuffix(word, "\n")
		if len(word) == 0 {
			fmt.Println("empty word!")
			continue
		}

		// check for substring in urls and titles
		for _, v := range hashMap {
			for url, title := range v {

				if strings.Contains(url, word) || strings.Contains(title, word) {
					fmt.Printf("%s - %s\n", url, title)
				}
			}
		}
	}
}