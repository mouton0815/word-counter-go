package main

func main() {
    numWorkers := 2 // TODO: Make number of workers configurable

    // TODO: Change size of all queues for experimenting
    pathQueue := make(chan string)
    wordQueue := make(chan string)
    doneQueue := make(chan bool, numWorkers)

    worker := NewWorker(pathQueue, wordQueue, doneQueue)
    for i := 0; i < numWorkers; i++ {
        go worker.Work(i)
    }

    terminator := newTerminator(numWorkers, wordQueue, doneQueue)
    go terminator.Terminate()

    gatherer := NewGatherer(pathQueue)
    go gatherer.Gather()

    counter := NewWordCounter(wordQueue)
    counter.Count()
    counter.Print()
}
