package test

import (
	"fmt"
	"sort"
	"testing"
	"word-counter-go/main/internal"
)

func TestWorkerPoolNoInput(t *testing.T) {
	RunWorkerPoolAndVerify(t, 0)
}

func TestWorkerPoolWithInput(t *testing.T) {
	RunWorkerPoolAndVerify(t, 100)
}

func RunWorkerPoolAndVerify(t *testing.T, pathCount int) {
	pathQueue := make(chan string, pathCount)
	refSlice := make([]string, 0, pathCount)
	for i := 0; i < pathCount; i++ {
		path := fmt.Sprintf("%02d", i)
		refSlice = append(refSlice, path)
		pathQueue <- path
	}
	close(pathQueue)

	fileReader := NewFileReaderMock(pathCount)
	workerPool := internal.NewWorkerPool(4, pathQueue, &fileReader)
	workerPool.Work()

	sort.Strings(fileReader.result) // Results may arrive unordered
	CompareStringSlices(t, fileReader.result, refSlice)
}
