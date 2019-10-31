package test

import (
    "log"
    "reflect"
    "testing"
    "word-counter-go/internal"
)

type FileReaderMock struct {
    result []string
}

func (r *FileReaderMock) Read(path string) {
    r.result = append(r.result, path)
}

func TestWorkerEmpty(t *testing.T) {
    WorkAndVerify(t)
}

func TestWorkerNormal(t *testing.T) {
    WorkAndVerify(t, "foo", "bar", "baz")
}

func WorkAndVerify(t *testing.T, paths ...string) {
    pathQueue := make(chan string)
    fileReader := &FileReaderMock{result: make([]string, 0, len(paths))}
    readyQueue := make(chan bool)
    worker := internal.NewWorker(pathQueue, fileReader, readyQueue)
    go worker.Work(1)

    for _, path := range paths {
        pathQueue <- path
    }
    close(pathQueue)

    // Block until worker is ready
    <- readyQueue

    refArray := CreateRefArray(paths)
    if !reflect.DeepEqual(fileReader.result, refArray) {
        log.Printf("'%v' <-> '%v'\n", fileReader.result, refArray)
        t.Fail()
    }
}
