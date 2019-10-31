package test

import (
    "log"
    "reflect"
    "testing"
    "word-counter-go/internal"
)

func TestTokenizerEmpty(t *testing.T) {
    TokenizeAndVerify(t, "  ")
}

func TestTokenizerSimple(t *testing.T) {
    TokenizeAndVerify(t, "foo", "foo")
}

func TestTokenizerTrim(t *testing.T) {
    TokenizeAndVerify(t, "  foo* ", "foo")
}

func TestTokenizerMulti(t *testing.T) {
    TokenizeAndVerify(t, "Foo bAr  baZ", "foo", "bar", "baz")
}

func TestTokenizerSpecialSigns(t *testing.T) {
    TokenizeAndVerify(t, "Foo-~bar #!  baz", "foo", "bar", "baz")
}

func TestTokenizerUnderscore(t *testing.T) {
    TokenizeAndVerify(t, "foo_bar__baz", "foo_bar__baz")
}

func TestTokenizerUnicodeGerman(t *testing.T) {
    TokenizeAndVerify(t, "Fö bär baß", "fö", "bär", "baß")
}

func TestTokenizerUnicodeRussian(t *testing.T) {
    TokenizeAndVerify(t, "Раз, два три!", "раз", "два", "три")
}

func TestTokenizerUnicodeCzech(t *testing.T) {
    TokenizeAndVerify(t, "Jedna, dva tři čtyři pět!", "jedna", "dva", "tři", "čtyři", "pět")
}

func TestIgnoreNumbers(t *testing.T) {
    TokenizeAndVerify(t, "foo 123 bar456baz", "foo", "bar", "baz")
}

func Tokenize(text string) []string {
    queue := make(chan string, 10)
    tokenizer := internal.NewTokenizer(queue)
    tokenizer.Tokenize(text)
    close(queue)

    words := make([]string, 0, 10)
    for word := range queue {
        words = append(words, word)
    }
    return words
}

func CreateRefArray(words []string) []string {
    ref := make([]string, 0, len(words))
    for _, word := range words {
        ref = append(ref, word)
    }
    return ref
}

func TokenizeAndVerify(t *testing.T, text string, refWords ...string) {
    words := Tokenize(text)
    ref := CreateRefArray(refWords)
    if !reflect.DeepEqual(words, ref) {
        log.Printf("'%v' <-> '%v'\n", words, ref)
        t.Fail()
    }
}
