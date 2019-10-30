package main

import (
    "runtime"
    "word-counter-go/internal"
)

func main() {
    // Reserve one CPU for the path collector and one for the word counter (this thread)
    numWorkers := runtime.NumCPU() - 2
    if numWorkers < 1 { numWorkers = 1 }

    pathQueue := make(chan string, 100)
    wordQueue := make(chan string, 1000)

    workerPool := internal.NewWorkerPool(numWorkers, pathQueue, wordQueue)
    go workerPool.Work()

    collector := internal.NewPathCollector(pathQueue)
    go collector.Collect("testdata")

    counter := internal.NewWordCounter(wordQueue)
    counter.Count()
    counter.Print()
}
