package internal

import (
    "fmt"
    "reflect"
    "testing"
)

func TestEmpty(t *testing.T) {
    AssertEqual(t, "  ")
}

func TestSimple(t *testing.T) {
    AssertEqual(t, "foo", "foo")
}

func TestTrim(t *testing.T) {
    AssertEqual(t, "  foo* ", "foo")
}

func TestMulti(t *testing.T) {
    AssertEqual(t, "Foo bAr  baZ", "foo", "bar", "baz")
}

func TestSpecialSigns(t *testing.T) {
    AssertEqual(t, "Foo-~bar #!  baz", "foo", "bar", "baz")
}

func TestUnderscore(t *testing.T) {
    AssertEqual(t, "foo_bar__baz", "foo_bar__baz")
}

func TestUnicode1(t *testing.T) {
    AssertEqual(t, "Fö bär baß", "fö", "bär", "baß")
}

func TestUnicode2(t *testing.T) {
    AssertEqual(t, "Раз, два три!", "раз", "два", "три")
}

func TestUnicode3(t *testing.T) {
    AssertEqual(t, "Jedna, dva tři čtyři pět!", "jedna", "dva", "tři", "čtyři", "pět")
}

func TestIgnoreNumbers(t *testing.T) {
    AssertEqual(t, "foo 123 bar456baz", "foo", "bar", "baz")
}

func Tokenize(text string) []string {
    queue := make(chan string, 10)
    tokenizer := NewTokenizer(queue)
    tokenizer.Tokenize(text)
    close(queue)

    words := make([]string, 0, 10)
    for word := range queue {
        words = append(words, word)
    }
    return words
}

func CreateRef(words []string) []string {
    ref := make([]string, 0, len(words))
    for _, word := range words {
        ref = append(ref, word)
    }
    return ref
}

func AssertEqual(t *testing.T, text string, refWords ...string) {
    words := Tokenize(text)
    ref := CreateRef(refWords)
    if !reflect.DeepEqual(words, ref) {
        fmt.Printf("'%v' <-> '%v'\n", words, ref)
        t.Fail()
    }
}
