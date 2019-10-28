package main

import (
    "fmt"
    "sort"
)

type Counter interface {
    Count(wordQueue chan string)
    Print()
}

type CounterImpl struct {
    wordCounts map[string]int
}

func (c *CounterImpl) Init()  {
    c.wordCounts = make(map[string]int)
}

func (c *CounterImpl) Count(wordQueue chan string)  {
    for word := range wordQueue {
        c.wordCounts[word]++
    }
}

// TODO: Move sorting to extra file?
// Sorting helpers
type WordCount struct {
    word string
    count int
}

type ByWordCount []WordCount

func (a ByWordCount) Len() int {
    return len(a)
}

// TODO: Why does this work considering that "a" is passed by value?
func (a ByWordCount) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a ByWordCount) Less(i, j int) bool {
    return a[i].count > a[j].count // TODO: Secondary sorting by value (alphabetically increasing)
}

func (c *CounterImpl) Print() {
    array := make([]WordCount, 0, len(c.wordCounts))
    for k, v := range c.wordCounts {
        array = append(array, WordCount{word: k, count: v})
    }
    sort.Sort(ByWordCount(array))
    for _, v := range array {
        fmt.Printf("%3d - %s\n", v.count, v.word)
    }
}

// Factory function
func NewCounter() Counter {
    var c CounterImpl
    c.Init()
    return &c
}
