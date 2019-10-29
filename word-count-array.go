//
// Typedef of an array of words with their counts plus sorting helpers
//
package main

type WordCount struct {
    word string
    count int
}

type WordCountArray []WordCount

func (a WordCountArray) Len() int {
    return len(a)
}

// TODO: Why does this work considering that "a" is passed (aka "received") by value?
func (a WordCountArray) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a WordCountArray) Less(i, j int) bool {
    if a[i].count == a[j].count {
        return a[i].word < a[j].word
    }
    return a[i].count > a[j].count
}

