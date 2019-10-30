package main

import (
    "fmt"
    "os"
    "runtime"
    "word-counter-go/internal"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Printf("Syntax %s <folder>\n", os.Args[0])
        return
    }
    rootPath := os.Args[1]

    // Reserve one CPU for the path collector and one for the word counter (this thread)
    numWorkers := runtime.NumCPU() - 2
    if numWorkers < 1 { numWorkers = 1 }

    pathQueue := make(chan string, 100)
    wordQueue := make(chan string, 1000)

    workerPool := internal.NewWorkerPool(numWorkers, pathQueue, wordQueue)
    go workerPool.Work()

    collector := internal.NewPathCollector(pathQueue)
    go collector.Collect(rootPath)

    counter := internal.NewWordCounter(wordQueue)
    counter.Count()
    counter.Print()
}
