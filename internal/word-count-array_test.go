package internal

import (
    "reflect"
    "sort"
    "testing"
)

func TestSortingEmpty(t *testing.T) {
    array := make(WordCountArray, 0, 0)
    sort.Sort(array)

    refArray := make(WordCountArray, 0, 0)
    if !reflect.DeepEqual(array, refArray) { t.Fail() }
}

func TestSortingSimple(t *testing.T) {
    foo := WordCount{word: "foo", count: 1}
    bar := WordCount{word: "bar", count: 2}

    array := make(WordCountArray, 0, 2)
    array = append(array, foo, bar)
    sort.Sort(array)

    refArray := make(WordCountArray, 0, 2)
    refArray = append(refArray, bar, foo)

    if !reflect.DeepEqual(array, refArray) { t.Fail() }
}

func TestSortingSecondary(t *testing.T) {
    foo := WordCount{word: "foo", count: 2}
    bar := WordCount{word: "bar", count: 1}
    baz := WordCount{word: "baz", count: 2}

    array := make(WordCountArray, 0, 3)
    array = append(array, foo, bar, baz)
    sort.Sort(array)

    refArray := make(WordCountArray, 0, 3)
    refArray = append(refArray, baz, foo, bar)
    if !reflect.DeepEqual(array, refArray) { t.Fail() }
}
