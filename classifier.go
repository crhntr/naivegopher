package naivegopher

import (
	"bufio"
	"io"
	"sync"
	"sync/atomic"
)

const minimumProbability = 0.000000001

type Classifier struct {
	*sync.RWMutex
	Categories []Category
	Learned    int
	Seen       int64
}

func NewClassifier() *Classifier {
	return &Classifier{
		RWMutex:    &sync.RWMutex{},
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
	c.RLock()
	defer c.RUnlock()

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
	c.RLock()
	defer c.RUnlock()

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

func (c Classifier) CategoryNames() []string {
	c.RLock()
	defer c.RUnlock()

	names := []string{}
	for _, category := range c.Categories {
		names = append(names, category.Name)
	}
	return names
}
