package naivegopher

type Category struct {
	Name            string
	WordFrequencies map[string]float64
	Total           int
}

// GetWordProbability returns P(W|C_j):
// the probability of seeing a particular word W in a document of this class.
func (c Category) GetWordProbability(word string) float64 {
	value, ok := c.WordFrequencies[word]
	if !ok {
		return defaultProabability
	}
	return float64(value) / float64(c.Total)
}

// GetWordsProbability returns P(D|C_j): the probability of seeing
// this set of words in a document of this class.
func (c Category) GetWordsProbability(words []string) float64 {
	p := float64(1.0)
	for _, word := range words {
		p *= c.GetWordProbability(word)
	}
	return p
}
