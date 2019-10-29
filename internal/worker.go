package internal

import (
    "fmt"
    "strings"
)

type Worker interface {
    Work(id int)
}

type WorkerImpl struct {
    pathQueue chan string
    wordQueue chan string
    doneQueue chan bool
}

func (w WorkerImpl) Work(id int) {
    fmt.Printf("Worker %d starts\n", id)
    for path := range w.pathQueue {
        fmt.Printf("Worker %d treats '%s'\n", id, path)
        wordArray := strings.Fields(path)
        for _, word := range wordArray {
            w.wordQueue <- word
        }
    }
    fmt.Printf("Worker %d leaves\n", id)
    w.doneQueue <- true
}

// Factory function
func NewWorker(pathQueue, wordQueue chan string, doneQueue chan bool) Worker {
    return WorkerImpl{ pathQueue: pathQueue, wordQueue: wordQueue, doneQueue: doneQueue }
}
