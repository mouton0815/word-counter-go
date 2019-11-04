package internal

type WorkerPool interface {
    Work()
}

type WorkerPoolImpl struct {
    numWorkers int
    pathQueue chan string
    fileReader FileReader
}

func (w WorkerPoolImpl) Work() {
    readyQueue := make(chan bool, w.numWorkers)
    worker := NewWorker(w.pathQueue, w.fileReader, readyQueue)

    // Spawn workers ...
    for i := 0; i < w.numWorkers; i++ {
        go worker.Work(i)
    }

    // ... and wait for their termination
    workerCount := w.numWorkers
    for range readyQueue {
        workerCount--
        if workerCount == 0 {
            w.fileReader.Close()
            close(readyQueue)
            return
        }
    }
}

// Factory function
func NewWorkerPool(numWorkers int, pathQueue chan string, fileReader FileReader) WorkerPool {
    return WorkerPoolImpl{ numWorkers, pathQueue, fileReader }
}
