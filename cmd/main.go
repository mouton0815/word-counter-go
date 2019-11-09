package main

import (
    "fmt"
    "log"
    "os"
    "runtime"
    "time"
    "word-counter-go/internal"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Printf("Syntax %s <folder>\n", os.Args[0])
        return
    }
    rootPath := os.Args[1]

    // Measure time including setup
    startTime := time.Now()

    // Reserve one CPU for the path collector and one for the word counter (this thread)
    numWorkers := runtime.NumCPU() - 2
    if numWorkers < 1 { numWorkers = 1 }

    pathQueue := make(chan string, 1000)
    wordQueue := make(chan string, 1000)

    fileReader := internal.NewFileReader(internal.NewTokenizer(wordQueue))
    workerPool := internal.NewWorkerPool(numWorkers, pathQueue, fileReader)
    go workerPool.Work()

    collector := internal.NewPathCollector(pathQueue)
    go collector.Collect(rootPath)

    counter := internal.NewWordCounter(wordQueue)
    wordCounts := counter.Count()
    for _, item := range wordCounts {
        fmt.Printf("%3d - %s\n", item.Count, item.Word)
    }

    elapsed := time.Since(startTime)
    log.Printf("Elapsed time %s\n", elapsed)
}
