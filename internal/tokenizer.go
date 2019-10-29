package internal

import (
    "regexp"
    "strings"
)

type Tokenizer interface {
    Tokenize(text string)
}

type TokenizerImpl struct {
    wordQueue chan string
}

func (t TokenizerImpl) Tokenize(text string) {
    pattern := regexp.MustCompile("[\\p{L}_]+")
    wordArray := pattern.FindAllString(text, -1)
    for _, word := range wordArray {
        t.wordQueue <- strings.ToLower(word)
    }
}

// Factory function
func NewTokenizer(wordQueue chan string) Tokenizer {
    return TokenizerImpl{ wordQueue: wordQueue }
}
