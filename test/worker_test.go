package test

import (
    "testing"
    "word-counter-go/internal"
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
    readyQueue := make(chan bool)
    worker := internal.NewWorker(pathQueue, &fileReader, readyQueue)
    go worker.Work(1)

    for _, path := range paths {
        pathQueue <- path
    }
    close(pathQueue)

    // Block until worker is ready
    <- readyQueue

    CompareStringSlices(t, fileReader.result, paths)
}
