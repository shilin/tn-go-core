package index

import (
	"sort"
	"strings"
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

type IndexedWords map[string][]int

func Build(hm map[string]string) (ById, IndexedWords, error) {
	indexedWords := IndexedWords{}
	docs := ById{}

	i := 0
	for url, title := range hm {
		newDoc := Doc{
			Id:    i,
			URL:   url,
			Title: title,
		}
		words := strings.Split(title, " ")
		words = append(words, url)

		for _, word := range words {
			indexedWords[word] = append(indexedWords[word], i)
		}

		i++
		docs = append(docs, newDoc)
	}
	sort.Sort(docs)
	return docs, indexedWords, nil
}
