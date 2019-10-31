package internal

import (
    "log"
    "sort"
)

type WordCountMap map[string]int

type WordCounter interface {
    Count() WordCountArray
}

type WordCounterImpl struct {
    wordQueue chan string
}

func (c WordCounterImpl) Count() WordCountArray {
    wordCounts := make(WordCountMap)
    for word := range c.wordQueue {
        wordCounts[word]++
    }
    log.Println("Counter leaves")
    return ToSortedArray(wordCounts)
}

func ToSortedArray(wordCounts WordCountMap) WordCountArray {
    array := make(WordCountArray, 0, len(wordCounts))
    for k, v := range wordCounts {
        array = append(array, WordCount{Word: k, Count: v})
    }
    sort.Sort(array)
    return array
}

// Factory function
func NewWordCounter(wordQueue chan string) WordCounter {
    return WordCounterImpl{ wordQueue: wordQueue }
}
