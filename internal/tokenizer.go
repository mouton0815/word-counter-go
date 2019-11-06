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
}

func (t TokenizerImpl) Tokenize(text string) {
    pattern := regexp.MustCompile("[\\p{L}_]+") // TODO: Compile regexp in constructor
    wordSlice := pattern.FindAllString(text, -1)
    for _, word := range wordSlice {
        t.wordQueue <- strings.ToLower(word)
    }
}

func (t TokenizerImpl) Close() {
    close(t.wordQueue)
}

// Factory function
func NewTokenizer(wordQueue chan string) Tokenizer {
    return TokenizerImpl{ wordQueue }
}
