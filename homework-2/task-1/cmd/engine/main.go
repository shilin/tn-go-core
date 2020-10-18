package main

import (
	"fmt"
	"log"

	"go-core.course/homework-2/task-1/pkg/spider"
)

func main() {
	sites := []string{"https://yandex.ru", "https://ya.ru"}
	depth := 2

	// result is a map of sites mapped to url to page title
	hashMap := make(map[string]map[string]string)

	for _, site := range sites {
		data, err := spider.Scan(site, depth)

		if err != nil {
			log.Println(err)
		}
		hashMap[site] = data
	}

	fmt.Println(hashMap)
}
