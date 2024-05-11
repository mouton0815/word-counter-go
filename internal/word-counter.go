package internal

import (
	"log"
	"sort"
)

type WordCountMap map[string]int

type WordCounter interface {
	Count() WordCountSlice
}

type WordCounterImpl struct {
	wordQueue chan string
}

func (c WordCounterImpl) Count() WordCountSlice {
	wordCounts := make(WordCountMap)
	for word := range c.wordQueue {
		wordCounts[word]++
	}
	log.Println("Counter leaves")
	return ToSortedSlice(wordCounts)
}

func ToSortedSlice(wordCounts WordCountMap) WordCountSlice {
	slice := make(WordCountSlice, 0, len(wordCounts))
	for k, v := range wordCounts {
		slice = append(slice, WordCount{Word: k, Count: v})
	}
	sort.Sort(slice)
	return slice
}

// Factory function
func NewWordCounter(wordQueue chan string) WordCounter {
	return WordCounterImpl{wordQueue: wordQueue}
}
