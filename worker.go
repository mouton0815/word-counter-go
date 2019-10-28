package main

import (
    "fmt"
    "strings"
)

type Worker interface {
    Work(id int, pathQueue chan string, wordQueue chan string, doneQueue chan bool)
}

type WorkerImpl struct {
    // TODO: Empty class
}

func (w *WorkerImpl) Work(id int, pathQueue chan string, wordQueue chan string, doneQueue chan bool) {
    fmt.Printf("Worker %d starts\n", id)
    for path := range pathQueue {
        fmt.Printf("Worker %d treats '%s'\n", id, path)
        wordArray := strings.Fields(path)
        for _, word := range wordArray {
            wordQueue <- word
        }
    }
    fmt.Printf("Worker %d leaves\n", id)
    doneQueue <- true
}

// Factory function
func NewWorker() Worker {
    return &WorkerImpl{}
}
