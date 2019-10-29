package main

type Terminator interface {
    Terminate()
}

type TerminatorImpl struct {
    numWorkers int
    wordQueue chan string
    doneQueue chan bool
}

func (t TerminatorImpl) Terminate() {
    workerCount := t.numWorkers
    for range t.doneQueue {
        workerCount--
        if workerCount == 0 {
            close(t.wordQueue)
            close(t.doneQueue)
            return
        }
    }
}

// Factory function
func newTerminator(numWorkers int, wordQueue chan string, doneQueue chan bool) Terminator {
    return TerminatorImpl{ numWorkers: numWorkers, wordQueue: wordQueue, doneQueue: doneQueue }
}

