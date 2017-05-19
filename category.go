package naivegopher

import "sync"

type Category struct {
	*sync.Mutex
	Name            string
	WordFrequencies map[string]int
	TotalWordCount  int
}

// GetWordProbability returns P(W|C_j):
// the probability of seeing a particular word W in a document of this class.
func (c Category) GetWordProbability(word string) float64 {
	c.Lock()
	defer c.Unlock()

	if value, ok := c.WordFrequencies[word]; ok {
		return float64(value) / float64(c.TotalWordCount)
	}
	return minimumProbability
}

// GetWordsProbability returns P(D|C_j): the probability of seeing
// this set (unique set) of words in a document in this category.
func (c Category) GetWordsProbability(words []string) float64 {
	c.Lock()
	defer c.Unlock()

	p := float64(1.0)
	filteredWords := words[:0]
	for i, word := range words {
		if uniqueInSlice(word, words[:i]) {
			filteredWords = append(filteredWords, word)
		}
	}
	for _, word := range filteredWords {
		p *= c.GetWordProbability(word)
	}
	if p < minimumProbability {
		return minimumProbability
	}
	return p
}

func uniqueInSlice(str string, slice []string) bool {
	count := 0
	for _, s := range slice {
		if str == s {
			if count++; count > 1 {
				return false
			}
		}
	}
	return true
}

func (category Category) String() string {
	return category.Name
}
