package index

import (
	"sort"
	"strings"
)

// Doc is a struct for representing page with URL and Title
type Doc struct {
	ID    int
	URL   string
	Title string
}

// byID is a type to implement esenting page with URL and Title
type byID []Doc

var docs byID

var nextID int

func (d byID) Len() int {
	return len(d)
}
func (d byID) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d byID) Less(i, j int) bool {
	return d[i].ID < d[j].ID
}

// Build produces slice of Docs and inverted index for words
func Build(hm map[string]string) (byID, map[string][]int) {
	indexedWords := map[string][]int{}

	for url, title := range hm {
		// new Doc will have id of the last Doc in slice plus one
		if docs.Len() > 0 {
			nextID = docs[docs.Len()-1].ID + 1
		}
		newDoc := Doc{
			ID:    nextID,
			URL:   url,
			Title: title,
		}
		words := strings.Split(title, " ")
		words = append(words, url)

		for _, word := range words {
			indexedWords[word] = appendID(indexedWords[word], nextID)
		}

		docs = append(docs, newDoc)
	}
	sort.Sort(docs)
	return docs, indexedWords
}

func appendID(nums []int, id int) []int {
	for _, v := range nums {
		if v == id {
			return nums
		} 
	}
	return append(nums, id)
}
