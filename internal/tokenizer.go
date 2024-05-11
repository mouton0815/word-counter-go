package internal

import (
	"regexp"
	"strings"
)

type Tokenizer interface {
	Tokenize(text string)
	Close()
}

type TokenizerImpl struct {
	wordQueue chan string
	pattern   *regexp.Regexp
}

func (t TokenizerImpl) Tokenize(text string) {
	wordSlice := t.pattern.FindAllString(text, -1)
	for _, word := range wordSlice {
		t.wordQueue <- strings.ToLower(word)
	}
}

func (t TokenizerImpl) Close() {
	close(t.wordQueue)
}

// Factory function
func NewTokenizer(wordQueue chan string) Tokenizer {
	return TokenizerImpl{wordQueue: wordQueue, pattern: regexp.MustCompile("[\\p{L}_]+")}
}
