package naivegopher

import (
	"bytes"
	"io"
)

type BytesReader interface {
	ReadBytes(delim byte) ([]byte, error)
}

func nextWord(sr BytesReader) string {
	for {
		wordCandidate, err := sr.ReadBytes(' ')
		if err != nil {
			if err == io.EOF {
				return ""
			}
			panic(err)
		}

		wordCandidate = trimWord(wordCandidate)
		if !skipWord(wordCandidate) {
			return string(wordCandidate)
		}
	}
}
func trimWord(wordCandidate []byte) []byte {
	wordCandidate = bytes.TrimSpace(wordCandidate)
	for _, cutset := range cutsets {
		wordCandidate = bytes.TrimPrefix(wordCandidate, cutset)
		wordCandidate = bytes.TrimSuffix(wordCandidate, cutset)
	}
	wordCandidate = bytes.Trim(wordCandidate, ".,:;!&#`\"'[]{}\\+-=_()%%^<>")

	return wordCandidate
}

func skipWord(word []byte) bool {
	for _, skip := range skipWords {
		if bytes.Equal(skip, word) {
			return true
		}
	}
	return false
}

var skipWords = [][]byte{
	[]byte("the"),
	[]byte("them"),
	[]byte("this"),
	[]byte("i"),
	[]byte("a"),
	[]byte("all"),
	[]byte("or"),
	[]byte("it"),
	[]byte("was"),
	[]byte("an"),
	[]byte("and"),
	[]byte("are"),
	[]byte("as"),
	[]byte("which"),
	[]byte("be"),
	[]byte("been"),
	[]byte("but"),
	[]byte("by"),
	[]byte("may"),
	[]byte("for"),
	[]byte("did"),
	[]byte("we"),
	// []byte("methods"),
	// []byte("conclusion"),
	[]byte("about"),
	[]byte("has"),
	[]byte("from"),
	[]byte("have"),
	[]byte("in"),
	[]byte("more"),
	[]byte("there"),
	[]byte("our"),
	[]byte("to"),
	[]byte("of"),
	[]byte("that"),
	[]byte("these"),
	[]byte("being"),
	[]byte("on"),
	[]byte("than"),
	[]byte("with"),
	[]byte("is"),
	[]byte("its"),
	[]byte("<p>"),
	[]byte("</p>"),
	[]byte("do"),
	[]byte("not"),
	[]byte("also"),
	[]byte("one"),
	[]byte("two"),
	[]byte("three"),
	[]byte("four"),
	[]byte("five"),
	[]byte("six"),
	[]byte("seven"),
	[]byte("eight"),
	[]byte("nine"),
	[]byte("ten"),
	[]byte("eleven"),
	[]byte("twelve"),
	[]byte("thirteen"),
	[]byte("fourtieen"),
	[]byte("fifteen"),
	[]byte("sixteen"),
	[]byte("seventeen"),
	[]byte("eighteen"),
	[]byte("nineteen"),
	[]byte("most"),
	[]byte("now"),
	[]byte("who"),
	[]byte("each"),
	[]byte("at"),
	[]byte("shows"),
	[]byte("can"),
	// []byte("amount"),
	[]byte("were"),
	[]byte("after"),
	[]byte("into"),
	[]byte("even"),
	[]byte("often"),
	[]byte("those"),
	[]byte("both"),
}
