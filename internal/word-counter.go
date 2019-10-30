package internal

import (
    "fmt"
    "log"
    "sort"
)

type WordCounter interface {
    Count()
    Print()
}

type WordCounterImpl struct {
    wordQueue chan string
    wordCounts map[string]int
}

func (c WordCounterImpl) Count()  {
    for word := range c.wordQueue {
        c.wordCounts[word]++
    }
    log.Println("Counter leaves")
}

func (c WordCounterImpl) Print() {
    array := make(WordCountArray, 0, len(c.wordCounts))
    for k, v := range c.wordCounts {
        array = append(array, WordCount{word: k, count: v})
    }
    sort.Sort(array)
    for _, v := range array {
        fmt.Printf("%3d - %s\n", v.count, v.word)
    }
}

// Factory function
func NewWordCounter(wordQueue chan string) WordCounter {
    return WordCounterImpl{ wordQueue: wordQueue, wordCounts: make(map[string]int) }
}
