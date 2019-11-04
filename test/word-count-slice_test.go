package test

import (
    "sort"
    "testing"
    "word-counter-go/internal"
)

func TestWordSliceSortingEmpty(t *testing.T) {
    slice := internal.WordCountSlice{}
    refSlice := internal.WordCountSlice{}
    SortAndVerify(t, slice, refSlice)
}

func TestWordSliceSortingSimple(t *testing.T) {
    foo := internal.WordCount{"foo", 1}
    bar := internal.WordCount{"bar", 2}

    slice := internal.WordCountSlice{foo, bar}
    refSlice := internal.WordCountSlice{bar, foo}
    SortAndVerify(t, slice, refSlice)
}

func TestWordSliceSortingSecondary(t *testing.T) {
    foo := internal.WordCount{"foo", 2}
    bar := internal.WordCount{"bar", 1}
    baz := internal.WordCount{"baz", 2}

    slice := internal.WordCountSlice{foo, bar, baz}
    refSlice := internal.WordCountSlice{baz, foo, bar}
    SortAndVerify(t, slice, refSlice)
}

func TestWordSliceSortingUnicode(t *testing.T) {
    foo := internal.WordCount{"über", 1}
    bar := internal.WordCount{"Zuse", 1}
    baz := internal.WordCount{"Ödem", 1}

    slice := internal.WordCountSlice{foo, bar, baz}
    refSlice := internal.WordCountSlice{baz, foo, bar}
    SortAndVerify(t, slice, refSlice)
}

func SortAndVerify(t *testing.T, slice internal.WordCountSlice, refSlice internal.WordCountSlice) {
    sort.Sort(slice)
    CompareWordCountSlices(t, slice, refSlice)
}
