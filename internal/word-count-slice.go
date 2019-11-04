//
// Typedef of a slice of words with their counts plus sorting helpers
//
package internal

import (
    "golang.org/x/text/collate"
    "golang.org/x/text/language"
)

type WordCount struct {
    Word string
    Count int
}

type WordCountSlice []WordCount

var wordCollator = collate.New(language.English, collate.Loose)

func (a WordCountSlice) Len() int {
    return len(a)
}

// TODO: Why does this work considering that "a" is passed (aka "received") by value?
func (a WordCountSlice) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a WordCountSlice) Less(i, j int) bool {
    if a[i].Count == a[j].Count {
        return wordCollator.CompareString(a[i].Word, a[j].Word) == -1
    }
    return a[i].Count > a[j].Count // Descending order
}
