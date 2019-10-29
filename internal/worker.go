package internal

import (
    "fmt"
)

type Worker interface {
    Work(id int)
}

type WorkerImpl struct {
    pathQueue chan string
    fileReader FileReader
    doneQueue chan bool
}

func (w WorkerImpl) Work(id int) {
    fmt.Printf("Worker %d starts\n", id)
    for path := range w.pathQueue {
        fmt.Printf("Worker %d reads '%s'\n", id, path)
        w.fileReader.Read(path)
   }
    fmt.Printf("Worker %d leaves\n", id)
    w.doneQueue <- true
}

// Factory function
func NewWorker(pathQueue chan string, fileReader FileReader, doneQueue chan bool) Worker {
    return WorkerImpl{ pathQueue: pathQueue, fileReader: fileReader, doneQueue: doneQueue }
}
