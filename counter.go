package main

import "fmt"

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

func (c *CounterImpl) Print() {
    fmt.Printf("====================\n")
    for k, v := range c.wordCounts {
        fmt.Printf("%3d - %s\n", v, k)
    }
    fmt.Printf("====================\n")
}

// Factory function
func NewCounter() Counter {
    var c CounterImpl
    c.Init()
    return &c
}
