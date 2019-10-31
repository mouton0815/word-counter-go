package internal

import (
    "log"
    "reflect"
    "sort"
    "testing"
)

func TestWordArraySortingEmpty(t *testing.T) {
    array := CreateWordCountArray()
    refArray := CreateWordCountArray()
    SortAndVerify(t, array, refArray)
}

func TestWordArraySortingSimple(t *testing.T) {
    foo := WordCount{"foo", 1}
    bar := WordCount{"bar", 2}

    array := CreateWordCountArray(foo, bar)
    refArray := CreateWordCountArray(bar, foo)
    SortAndVerify(t, array, refArray)
}

func TestWordArraySortingSecondary(t *testing.T) {
    foo := WordCount{"foo", 2}
    bar := WordCount{"bar", 1}
    baz := WordCount{"baz", 2}

    array := CreateWordCountArray(foo, bar, baz)
    refArray := CreateWordCountArray(baz, foo, bar)
    SortAndVerify(t, array, refArray)
}

func TestWordArraySortingUnicode(t *testing.T) {
    foo := WordCount{"über", 1}
    bar := WordCount{"Zuse", 1}
    baz := WordCount{"Ödem", 1}

    array := CreateWordCountArray(foo, bar, baz)
    refArray := CreateWordCountArray(baz, foo, bar)
    SortAndVerify(t, array, refArray)
}

func CreateWordCountArray(wordCounts ...WordCount) WordCountArray {
    array := make(WordCountArray, 0, len(wordCounts))
    for _, wordCount := range wordCounts {
        array = append(array, wordCount)
    }
    return array
}

func SortAndVerify(t *testing.T, array WordCountArray, refArray WordCountArray) {
    sort.Sort(array)
    if !reflect.DeepEqual(array, refArray) {
        log.Printf("'%v' <-> '%v'\n", array, refArray)
        t.Fail()
    }
}
