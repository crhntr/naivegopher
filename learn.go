package naivegopher

import (
	"bufio"
	"io"
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
	reader := bufio.NewReaderSize(r, bufferSize)
	category := classifier.FindOrInsert(categoryName)
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
