package test

import (
    "strings"
    "testing"
    "word-counter-go/internal"
)

func TestPathCollector(t *testing.T) {
    rootPath := "../data"

    pathQueue := make(chan string, 10)
    pathCollector := internal.NewPathCollector(pathQueue)
    pathCollector.Collect(rootPath)

    slice := make([]string, 0, 10)
    for path := range pathQueue {
        slice = append(slice, strings.Replace(path, "\\", "/", -1))
    }

    refSlice := []string{rootPath + "/file1.txt", rootPath + "/file2.txt", rootPath + "/subdir/file3.txt"}
    CompareStringSlices(t, slice, refSlice)
}
