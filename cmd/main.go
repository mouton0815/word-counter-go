package main

import "word-counter-go/internal"

func main() {
    numWorkers := 4 // TODO: Make number of workers configurable (by command line? by #cores?)

    // TODO: Change size of all queues for experimenting
    pathQueue := make(chan string)
    wordQueue := make(chan string)

    workerPool := internal.NewWorkerPool(numWorkers, pathQueue, wordQueue)
    go workerPool.Work()

    collector := internal.NewPathCollector(pathQueue)
    go collector.Collect("testdata")

    counter := internal.NewWordCounter(wordQueue)
    counter.Count()
    counter.Print()
}
