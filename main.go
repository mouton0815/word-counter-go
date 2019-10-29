package main

func main() {
    numWorkers := 2 // TODO: Make number of workers configurable (by command line? by #cores?)

    // TODO: Change size of all queues for experimenting
    pathQueue := make(chan string)
    wordQueue := make(chan string)

    workerPool := NewWorkerPool(numWorkers, pathQueue, wordQueue)
    go workerPool.Work()

    gatherer := NewGatherer(pathQueue)
    go gatherer.Gather()

    counter := NewWordCounter(wordQueue)
    counter.Count()
    counter.Print()
}
