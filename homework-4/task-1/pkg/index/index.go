package index

import (
	"fmt"
	"strings"
	"sort"
)

type Doc struct {
	Id    int
	URL   string
	Title string
}

type ById []Doc

func (d ById) Len() int {
	return len(d)
}
func (d ById) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d ById) Less(i, j int) bool {
	return d[i].Id < d[j].Id
}


func Build(hm map[string]string) (ById, map[string][]int, error) {
	indexedWords := map[string][]int{}

	// hm1 := map[string]string{
	// 	"https://yandex.ru/collections/":  "Яндекс.Коллекции",
	// 	"https://yandex.ru/all":           "Все сервисы Яндекса",
	// 	"https://yandex.ru/blog/company/": "Блог Яндекса",
	// }

	docs := ById{}

	i := 0
	for url, title := range hm {
		newDoc := Doc{
			Id: i,
			URL: url,
			Title: title,
		}
		words := strings.Split(title, " ")
		words = append(words, url)

		for _, word := range words {
			indexedWords[word] = append(indexedWords[word], i)
		}

		i++

		docs = append(docs, newDoc)
		// fmt.Printf("id - %d. url - %s, title - %s\n", newDoc.Id, newDoc.URL, newDoc.Title)
	}
		// fmt.Printf("source url - %s. url - %s, title - %s\n", source, url, title)
		fmt.Println(docs)

		sort.Sort(docs)
		// fmt.Println(docs)
	return docs, indexedWords, nil
}
