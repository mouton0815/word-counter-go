//
// Typedef of an array of words with their counts plus sorting helpers
//
package internal

type WordCount struct {
    Word string
    Count int
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
    if a[i].Count == a[j].Count {
        return a[i].Word < a[j].Word
    }
    return a[i].Count > a[j].Count
}

