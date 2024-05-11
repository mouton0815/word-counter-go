package internal

import (
	"log"
	"sync"
)

type Worker interface {
	Work(id int)
}

type WorkerImpl struct {
	pathQueue  chan string
	fileReader FileReader
	waitGroup  *sync.WaitGroup
}

func (w WorkerImpl) Work(id int) {
	log.Printf("Worker %d starts\n", id)
	defer w.waitGroup.Done()
	for path := range w.pathQueue {
		log.Printf("Worker %d reads '%s'\n", id, path)
		w.fileReader.Read(path)
	}
	log.Printf("Worker %d leaves\n", id)
}

// Factory function
func NewWorker(pathQueue chan string, fileReader FileReader, waitGroup *sync.WaitGroup) Worker {
	return WorkerImpl{pathQueue, fileReader, waitGroup}
}
