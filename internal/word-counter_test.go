package internal

import (
    "log"
    "reflect"
    "testing"
)

func TestWordCounterEmpty(t *testing.T) {
    wordQueue := CreateWordQueue()
    refArray := CreateWordCountRef()
    CountAndVerify(t, wordQueue, refArray)
}

func TestWordCounterSingle(t *testing.T) {
    wordQueue := CreateWordQueue("foo")
    refArray := CreateWordCountRef(WordCount{"foo", 1})
    CountAndVerify(t, wordQueue, refArray)
}

func TestWordCounterMultiple(t *testing.T) {
    wordQueue := CreateWordQueue("bar", "foo", "bar")
    refArray := CreateWordCountRef(WordCount{ "bar", 2}, WordCount{"foo", 1})
    CountAndVerify(t, wordQueue, refArray)
}

func CreateWordQueue(words ...string) chan string {
    wordQueue := make(chan string, len(words))
    for _, word := range words {
        wordQueue <- word
    }
    close(wordQueue)
    return wordQueue
}

func CreateWordCountRef(wordCounts ...WordCount) WordCountArray {
    array := make(WordCountArray, 0, len(wordCounts))
    for _, wordCount := range wordCounts {
        array = append(array, wordCount)
    }
    return array
}

func CountAndVerify(t *testing.T, wordQueue chan string, refArray WordCountArray) {
    counter := NewWordCounter(wordQueue)
    result := counter.Count()

    if !reflect.DeepEqual(result, refArray) {
        log.Printf("'%v' <-> '%v'\n", result, refArray)
        t.Fail()
    }
}


