package index

import (
	"strings"

	"go-core.course/pkg/btree"
)

// Doc is a struct for representing page with URL and Title
type Doc struct {
	ID    int
	URL   string
	Title string
}

// Identity method complies to Identifier interface
func (d *Doc) Identity() int {
	return d.ID
}

// URLstring method complies to Identifier interface
func (d *Doc) URLstring() string {
	return d.URL
}

// Titlestring method complies to Identifier interface
func (d *Doc) Titlestring() string {
	return d.Title
}

func generateID(i int, ids []int) (newI int, newID int, newIDs []int) {
	// cycle through slice of ascenidng natural numbers sequence
	// using remainder of division by three
	rem := i % 3
	i = i + 1
	idsLen := len(ids)
	pos := 0

	// initial step - take the center of the whole
	if rem == 0 {
		pos = idsLen / 2
	}

	// second step - take the center of the left half
	if rem == 1 {
		pos = idsLen / 4
	}

	// third step - take the center of the right half
	// next step after third is initial again...
	if rem == 2 {
		pos = 3 * idsLen / 4
	}

	id := ids[pos]
	ids = del(ids, pos)
	return i, id, ids
}

// Build produces binarytree of Docs and inverted index for words
func Build(hm map[string]string) (*btree.Node, map[string][]int) {
	indexedWords := map[string][]int{}

	docs := new(btree.Node)
	i := 0
	id := 0
	idsLen := len(hm)
	ids := make([]int, idsLen)
	for j := 0; j < idsLen; j++ {
		ids[j] = j
	}

	for url, title := range hm {
		i, id, ids = generateID(i, ids)

		newDoc := Doc{
			ID:    id,
			URL:   url,
			Title: title,
		}

		newNode := new(btree.Node)
		newNode.Identifier = &newDoc

		if docs.Identifier == nil {
			docs = newNode
		} else {
			docs.Insert(newNode)
		}

		words := strings.Fields(title)
		words = append(words, url)

		for _, word := range words {
			indexedWords[word] = appendID(indexedWords[word], id)
		}
	}

	return docs, indexedWords
}

func del(s []int, pos int) []int {
	copy(s[pos:], s[pos+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]
	return s
}

func appendID(nums []int, id int) []int {
	for _, v := range nums {
		if v == id {
			return nums
		}
	}
	return append(nums, id)
}
