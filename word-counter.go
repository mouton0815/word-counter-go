package main

import (
    "fmt"
    "sort"
)

type WordCounter interface {
    Count(wordQueue chan string)
    Print()
}

type WordCounterImpl struct {
    wordCounts map[string]int
}

func (c WordCounterImpl) Count(wordQueue chan string)  {
    for word := range wordQueue {
        c.wordCounts[word]++
    }
    fmt.Println("WordCounter leaves")
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
func NewWordCounter() WordCounter {
    return WordCounterImpl{ wordCounts: make(map[string]int) }
}
