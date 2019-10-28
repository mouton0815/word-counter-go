package main

type Terminator interface {
    Terminate(wordQueue chan string, doneQueue chan bool)
}

type TerminatorImpl struct {
    numWorkers int
}

func (t *TerminatorImpl) Terminate(wordQueue chan string, doneQueue chan bool) {
    workerCount := t.numWorkers
    for range doneQueue {
        workerCount--
        if workerCount == 0 {
            close(wordQueue)
            close(doneQueue)
            return
        }
    }
}

// Factory function
func newTerminator(numWorkers int) Terminator {
    return &TerminatorImpl{ numWorkers }
}

