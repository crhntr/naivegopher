package naivegopher

import (
	"bufio"
	"io"
	"sort"
	"strings"
	"sync/atomic"
)

const minimumProbability = 0.000000001

type Classifier struct {
	// *sync.RWMutex
	Categories []Category
	Learned    int
	Seen       int64
}

func NewClassifier() *Classifier {
	return &Classifier{
		// RWMutex:    &sync.RWMutex{},
		Categories: []Category{},
		Learned:    0,
		Seen:       0,
	}
}

type Scores []float64

func (s Scores) Max() int {
	return max([]float64(s))
}

func max(slc []float64) int {
	maxIndex := 0
	for ri, rv := range slc {
		if rv > slc[maxIndex] {
			maxIndex = ri
		}
	}
	return maxIndex
}

func (c *Classifier) ProbableCategoreies(r io.Reader) (scores Scores) {
	n := len(c.Categories)
	scores = make([]float64, n)

	if c.Learned < 1 {
		return // scores, categories
	}

	priors := c.PriorProbabilities()
	reader := bufio.NewReaderSize(r, bufferSize)

	for i, p := range priors {
		scores[i] = p
	}

	var sum float64 = 0
	// c is the sum of the logarithms
	// as outlined in the refresher
	for {
		word, done := nextWord(reader)
		if word == "" && done {
			break
		}
		for i, category := range c.Categories {

			p := category.GetWordProbability(word)
			scores[i] *= p

			sum += scores[i]
		}
	}

	for i := 0; i < n; i++ {
		scores[i] /= sum
	}

	atomic.AddInt64(&c.Seen, 1)
	return // scores, categories
}

// PriorProbabilities returns the prior probabilities for the
// classes provided -- P(C_j).
//
func (c *Classifier) PriorProbabilities() []float64 {
	// c.RLock()
	// defer c.RUnlock()

	n := len(c.Categories)
	priors := make([]float64, n, n)
	sum := 0

	for i, data := range c.Categories {
		total := data.TotalWordCount
		priors[i] = float64(total)
		sum += total
	}
	if sum != 0 {
		for i := 0; i < n; i++ {
			priors[i] /= float64(sum)
		}
	}
	return priors
}

func (c *Classifier) Len() int {
	// c.RLock()
	// defer c.RUnlock()
	return len(c.Categories)
}
func (c *Classifier) Swap(i, j int) {
	// c.Lock()
	// defer c.Unlock()
	c.Categories[i], c.Categories[j] = c.Categories[j], c.Categories[i]
}
func (c *Classifier) Less(i, j int) bool {
	// c.RLock()
	// defer c.RUnlock()
	return strings.Compare(c.Categories[i].Name, c.Categories[j].Name) < 0
}

// FindCategory returns the index of the category with a given name
func (c *Classifier) FindCategory(name string) int {
	// c.RLock()
	// defer c.RUnlock()
	for i, category := range c.Categories {
		if category.Name == name {
			return i
		}
	}
	return -1
}

// FindOrInsert searches for a Category with categoryName
// if it does not find one it inserts a new category in the
// correct ordered location
func (c *Classifier) FindOrInsert(categoryName string) *Category {
	// c.Lock()
	// defer c.Unlock()
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
		// Mutex:           &sync.Mutex{},
	}
	return &c.Categories[i]
}

func (c Classifier) CategoryNames() []string {
	// c.RLock()
	// defer c.Unlock()
	names := []string{}
	for _, category := range c.Categories {
		names = append(names, category.Name)
	}
	return names
}
