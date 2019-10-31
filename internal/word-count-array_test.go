package internal

import (
    "reflect"
    "sort"
    "testing"
)

func TestWordArraySortingEmpty(t *testing.T) {
    array := make(WordCountArray, 0, 0)
    sort.Sort(array)

    refArray := make(WordCountArray, 0, 0)
    if !reflect.DeepEqual(array, refArray) { t.Fail() }
}

func TestWordArraySortingSimple(t *testing.T) {
    foo := WordCount{Word: "foo", Count: 1}
    bar := WordCount{Word: "bar", Count: 2}

    array := make(WordCountArray, 0, 2)
    array = append(array, foo, bar)
    sort.Sort(array)

    refArray := make(WordCountArray, 0, 2)
    refArray = append(refArray, bar, foo)

    if !reflect.DeepEqual(array, refArray) { t.Fail() }
}

func TestWordArraySortingSecondary(t *testing.T) {
    foo := WordCount{Word: "foo", Count: 2}
    bar := WordCount{Word: "bar", Count: 1}
    baz := WordCount{Word: "baz", Count: 2}

    array := make(WordCountArray, 0, 3)
    array = append(array, foo, bar, baz)
    sort.Sort(array)

    refArray := make(WordCountArray, 0, 3)
    refArray = append(refArray, baz, foo, bar)
    if !reflect.DeepEqual(array, refArray) { t.Fail() }
}
