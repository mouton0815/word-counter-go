package test

import (
	"sync"
	"testing"
	"word-counter-go/main/internal"
)

func TestWorkerEmpty(t *testing.T) {
	WorkAndVerify(t)
}

func TestWorkerNormal(t *testing.T) {
	WorkAndVerify(t, "foo", "bar", "baz")
}

func WorkAndVerify(t *testing.T, paths ...string) {
	pathQueue := make(chan string)
	fileReader := NewFileReaderMock(len(paths))
	var waitGroup sync.WaitGroup
	worker := internal.NewWorker(pathQueue, &fileReader, &waitGroup)
	waitGroup.Add(1)
	go worker.Work(1)

	for _, path := range paths {
		pathQueue <- path
	}
	close(pathQueue)

	// Block until worker is ready
	waitGroup.Wait()

	CompareStringSlices(t, fileReader.result, paths)
}
