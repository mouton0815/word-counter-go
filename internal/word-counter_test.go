package internal

import (
    "log"
    "reflect"
    "testing"
)

func TestWordCounterEmpty(t *testing.T) {
    queue := CreateWordQueue()
    array := CreateWordCountArray()
    CountAndVerify(t, queue, array)
}

func TestWordCounterSingle(t *testing.T) {
    queue := CreateWordQueue("foo")
    array := CreateWordCountArray(WordCount{"foo", 1})
    CountAndVerify(t, queue, array)
}

func TestWordCounterMultiple(t *testing.T) {
    queue := CreateWordQueue("bar", "foo", "bar")
    array := CreateWordCountArray(WordCount{"bar", 2}, WordCount{"foo", 1})
    CountAndVerify(t, queue, array)
}

func CreateWordQueue(words ...string) chan string {
    wordQueue := make(chan string, len(words))
    for _, word := range words {
        wordQueue <- word
    }
    close(wordQueue)
    return wordQueue
}

func CountAndVerify(t *testing.T, wordQueue chan string, refArray WordCountArray) {
    counter := NewWordCounter(wordQueue)
    result := counter.Count()

    if !reflect.DeepEqual(result, refArray) {
        log.Printf("'%v' <-> '%v'\n", result, refArray)
        t.Fail()
    }
}
