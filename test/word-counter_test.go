package test

import (
	"testing"
	"word-counter-go/main/internal"
)

func TestWordCounterEmpty(t *testing.T) {
	queue := CreateWordQueue()
	slice := internal.WordCountSlice{}
	CountAndVerify(t, queue, slice)
}

func TestWordCounterSingle(t *testing.T) {
	queue := CreateWordQueue("foo")
	slice := internal.WordCountSlice{internal.WordCount{"foo", 1}}
	CountAndVerify(t, queue, slice)
}

func TestWordCounterMultiple(t *testing.T) {
	queue := CreateWordQueue("bar", "foo", "bar")
	slice := internal.WordCountSlice{internal.WordCount{"bar", 2}, internal.WordCount{"foo", 1}}
	CountAndVerify(t, queue, slice)
}

func CreateWordQueue(words ...string) chan string {
	wordQueue := make(chan string, len(words))
	for _, word := range words {
		wordQueue <- word
	}
	close(wordQueue)
	return wordQueue
}

func CountAndVerify(t *testing.T, wordQueue chan string, refSlice internal.WordCountSlice) {
	counter := internal.NewWordCounter(wordQueue)
	result := counter.Count()
	CompareWordCountSlices(t, result, refSlice)
}
