package internal

type WorkerPool interface {
    Work()
}

type WorkerPoolImpl struct {
    numWorkers int
    pathQueue chan string
    wordQueue chan string
}

func (w WorkerPoolImpl) Work() {
    doneQueue := make(chan bool, w.numWorkers)
    fileReader := NewFileReader(NewTokenizer(w.wordQueue))
    worker := NewWorker(w.pathQueue, fileReader, doneQueue)

    // Spawn workers ...
    for i := 0; i < w.numWorkers; i++ {
        go worker.Work(i)
    }

    // ... and wait for their termination
    workerCount := w.numWorkers
    for range doneQueue {
        workerCount--
        if workerCount == 0 {
            close(w.wordQueue)
            close(doneQueue)
            return
        }
    }
}

// Factory function
func NewWorkerPool(numWorkers int, pathQueue, wordQueue chan string) WorkerPool {
    return WorkerPoolImpl{ numWorkers: numWorkers, pathQueue: pathQueue, wordQueue: wordQueue }
}
