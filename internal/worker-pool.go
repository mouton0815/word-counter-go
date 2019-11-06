package internal

import "sync"

type WorkerPool interface {
    Work()
}

type WorkerPoolImpl struct {
    numWorkers int
    pathQueue chan string
    fileReader FileReader
}

func (w WorkerPoolImpl) Work() {
    var waitGroup sync.WaitGroup
    worker := NewWorker(w.pathQueue, w.fileReader, &waitGroup)

    // Spawn workers ...
    for i := 0; i < w.numWorkers; i++ {
        waitGroup.Add(1)
        go worker.Work(i)
    }

    // ... and wait for their termination
    waitGroup.Wait()
    w.fileReader.Close()
}

// Factory function
func NewWorkerPool(numWorkers int, pathQueue chan string, fileReader FileReader) WorkerPool {
    return WorkerPoolImpl{ numWorkers, pathQueue, fileReader }
}
