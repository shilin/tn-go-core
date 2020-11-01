package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sort"

	"go-core.course/homework-4/task-1/pkg/spider"
	// we can import a substitution for a spider pkg
	"go-core.course/homework-4/task-1/pkg/index"
	"go-core.course/homework-4/task-1/pkg/scan"
)

func main() {
	sites := []string{"https://yandex.ru", "https://ya.ru"}
	depth := 2

	s := new(spider.Spider)
	parsedSites := scanResults(s, sites, depth)
	docs, indexedWords, _ := index.Build(parsedSites)

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
		word = strings.ToLower(word)
		if len(word) == 0 {
			fmt.Println("empty word!")
			continue
		}

		resultNums := []int{}
		// check for substring in urls and titles
		for indexedWord, docNums := range indexedWords {
			if strings.Contains(strings.ToLower(indexedWord), word) {
				resultNums = append(resultNums, docNums...)
			}
		}
		fmt.Println(resultNums)

		for _, num := range resultNums {
			j := sort.Search(len(docs), func(i int) bool { return docs[i].Id == num })
			if j < len(docs)  {
				fmt.Printf("Слово встречается в %s - %s \n", docs[j].URL, docs[j].Title)
			}
		}
	}
}

func scanResults(s scan.Scanner, sites []string, depth int) map[string]string {
	// result is a map of urls to page title
	parsedSites := make(map[string]string)

	fmt.Println("Please wait, scanning sites...")

	for _, site := range sites {
		data, err := s.Scan(site, depth)
		if err != nil {
			log.Println(err)
			continue
		}
		// merge new data
		for k, v := range data {
			parsedSites[k] = v
		}
	}
	return parsedSites
}
