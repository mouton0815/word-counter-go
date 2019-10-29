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
    pattern := regexp.MustCompile(`[^\w]+`)
    wordArray := pattern.Split(text, -1)
    for _, word := range wordArray {
        if len(word) > 0 { // Ignore leading and trailing empty tokens
            t.wordQueue <- strings.ToLower(word)
        }
    }
}

// Factory function
func NewTokenizer(wordQueue chan string) Tokenizer {
    return TokenizerImpl{ wordQueue: wordQueue }
}
