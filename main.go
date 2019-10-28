package main

func main() {
    numWorkers := 2 // TODO: Make number of workers configurable

    // TODO: Decrease size of all queues for experimenting
    pathQueue := make(chan string)
    wordQueue := make(chan string)
    doneQueue := make(chan bool, numWorkers)

    worker := NewWorker()
    for i := 0; i < numWorkers; i++ {
        go worker.Work(i, pathQueue, wordQueue, doneQueue)
    }

    terminator := newTerminator(numWorkers)
    go terminator.Terminate(wordQueue, doneQueue)

    gatherer := NewGatherer()
    go gatherer.Gather(pathQueue)

    counter := NewCounter()
    counter.Count(wordQueue)
    counter.Print()
}
