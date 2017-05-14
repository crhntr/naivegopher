package naivegopher

import (
	"bufio"
	"io"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
)

const defaultProabability = 0.00000000001

type Classifier struct {
	sync.RWMutex
	Categories []Category
	Learned    int
	Seen       int64
}

func (c *Classifier) ProbableCategoreies(r io.Reader) ([]float64, []*Category) {
	n := len(c.Categories)
	scores := make([]float64, n)
	categories := make([]*Category, n)
	priors := c.PriorProbabilities()
	reader := bufio.NewReaderSize(r, bufferSize)

	sum := float64(0)
	for index, category := range c.Categories {
		// c is the sum of the logarithms
		// as outlined in the refresher
		score := priors[index]
		for {
			word := nextWord(reader)
			if word == "" {
				break
			}
			score *= category.GetWordProbability(word)
		}
		scores[index] = score
		sum += score
	}
	for i := 0; i < n; i++ {
		scores[i] /= sum
	}
	atomic.AddInt64(&c.Seen, 1)
	return scores, categories
}

// PriorProbabilities returns the prior probabilities for the
// classes provided -- P(C_j).
//
func (c *Classifier) PriorProbabilities() []float64 {
	c.RLock()
	defer c.RUnlock()
	n := len(c.Categories)
	priors := make([]float64, n, n)
	sum := 0

	i := 0
	for _, data := range c.Categories {
		total := data.Total
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
	c.RLock()
	defer c.RUnlock()
	return len(c.Categories)
}
func (c *Classifier) Swap(i, j int) {
	c.Lock()
	defer c.Unlock()
	c.Categories[i], c.Categories[j] = c.Categories[j], c.Categories[i]
}
func (c *Classifier) Less(i, j int) bool {
	c.RLock()
	defer c.RUnlock()
	return strings.Compare(c.Categories[i].Name, c.Categories[j].Name) < 0
}

// FindCategory returns the index of the category with a given name
func (c *Classifier) FindCategory(name string) int {
	c.RLock()
	defer c.RUnlock()
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
	c.Lock()
	defer c.Unlock()
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
		WordFrequencies: make(map[string]float64),
		Total:           0,
	}
	return &c.Categories[i]
}
