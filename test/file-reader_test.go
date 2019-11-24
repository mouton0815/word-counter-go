package test

import (
    "log"
    "testing"
    "word-counter-go/internal"
)

func TestFileReaderSuccess(t *testing.T) {
    wordQueue := make(chan string, 100)
    tokenizer := internal.NewTokenizer(wordQueue)
    fileReader := internal.NewFileReader(tokenizer)
    fileReader.Read("../data/file1.txt")
    fileReader.Close()

    slice := make([]string, 0, 100)
    for word := range wordQueue {
        slice = append(slice, word)
    }

    if len(slice) != 100 {
        log.Printf("Queue should have 100 items, but has %d", len(slice))
        t.Fail()
    }
    if slice[0] != "lorem" {
        log.Printf("First word should be 'loret', but is '%s'", slice[0])
        t.Fail()
    }
    if slice[len(slice) - 1] != "amet" {
        log.Printf("Last word should be 'amet', but is '%s'", slice[len(slice) - 1])
        t.Fail()
    }
}

func TestFileReaderNotFound(t *testing.T) {
    defer func() {
        r := recover()
        if r == nil {
            t.Errorf("The code did not panic")
        }
    }()

    wordQueue := make(chan string, 100)
    tokenizer := internal.NewTokenizer(wordQueue)
    fileReader := internal.NewFileReader(tokenizer)
    fileReader.Read("does/not.exist")
}
