package naivegopher

import (
	"bufio"
	"io"
	"sort"
	"strings"
	"sync"
)

const (
	bufferSize = 200
)

var (
	cutsets = [][]byte{
		[]byte("<p>"),
		[]byte("</p>"),
	}
)

// Learn will accept new training documents for
// supervised learning.
func (classifier *Classifier) Learn(categoryName string, r io.Reader) {
	classifier.Lock()
	defer classifier.Unlock()

	reader := bufio.NewReaderSize(r, bufferSize)
	category := classifier.findOrInsert(categoryName)
	for {
		word, done := nextWord(reader)
		if word == "" && done {
			break
		}
		category.WordFrequencies[word]++
		category.TotalWordCount++
	}
	classifier.Learned++
}

// findOrInsert searches for a Category with categoryName
// if it does not find one it inserts a new category in the
// correct ordered location
func (c *Classifier) findOrInsert(categoryName string) *Category {
	i := sort.Search(c.Len(), func(i int) bool {
		return strings.Compare(c.Categories[i].Name, categoryName) >= 0
	})
	if i < c.Len() && c.Categories[i].Name == categoryName {
		// categoryName is present at c.Categories[i]
		return &c.Categories[i]
	}

	// categoryName is not present in data,
	// but i is the index where it would be inserted.
	c.Categories = append(c.Categories, Category{})
	copy(c.Categories[i+1:], c.Categories[i:])
	c.Categories[i] = Category{
		Name:            categoryName,
		WordFrequencies: make(map[string]int),
		TotalWordCount:  0,
		Mutex:           &sync.Mutex{},
	}
	return &c.Categories[i]
}

func (c *Classifier) Len() int {
	return len(c.Categories)
}
func (c *Classifier) Swap(i, j int) {
	c.Categories[i], c.Categories[j] = c.Categories[j], c.Categories[i]
}
func (c *Classifier) Less(i, j int) bool {
	return strings.Compare(c.Categories[i].Name, c.Categories[j].Name) < 0
}
