package test

import (
    "log"
    "reflect"
    "sort"
    "testing"
    "word-counter-go/internal"
)

func TestWordArraySortingEmpty(t *testing.T) {
    array := CreateWordCountArray()
    refArray := CreateWordCountArray()
    SortAndVerify(t, array, refArray)
}

func TestWordArraySortingSimple(t *testing.T) {
    foo := internal.WordCount{"foo", 1}
    bar := internal.WordCount{"bar", 2}

    array := CreateWordCountArray(foo, bar)
    refArray := CreateWordCountArray(bar, foo)
    SortAndVerify(t, array, refArray)
}

func TestWordArraySortingSecondary(t *testing.T) {
    foo := internal.WordCount{"foo", 2}
    bar := internal.WordCount{"bar", 1}
    baz := internal.WordCount{"baz", 2}

    array := CreateWordCountArray(foo, bar, baz)
    refArray := CreateWordCountArray(baz, foo, bar)
    SortAndVerify(t, array, refArray)
}

func TestWordArraySortingUnicode(t *testing.T) {
    foo := internal.WordCount{"über", 1}
    bar := internal.WordCount{"Zuse", 1}
    baz := internal.WordCount{"Ödem", 1}

    array := CreateWordCountArray(foo, bar, baz)
    refArray := CreateWordCountArray(baz, foo, bar)
    SortAndVerify(t, array, refArray)
}

func CreateWordCountArray(wordCounts ...internal.WordCount) internal.WordCountArray {
    array := make(internal.WordCountArray, 0, len(wordCounts))
    for _, wordCount := range wordCounts {
        array = append(array, wordCount)
    }
    return array
}

func SortAndVerify(t *testing.T, array internal.WordCountArray, refArray internal.WordCountArray) {
    sort.Sort(array)
    if !reflect.DeepEqual(array, refArray) {
        log.Printf("'%v' <-> '%v'\n", array, refArray)
        t.Fail()
    }
}
