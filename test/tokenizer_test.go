package test

import (
	"testing"
	"word-counter-go/main/internal"
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

func TestTokenizerIgnoreNumbers(t *testing.T) {
	TokenizeAndVerify(t, "foo 123 bar456baz", "foo", "bar", "baz")
}

func Tokenize(text string) []string {
	queue := make(chan string, 10)
	tokenizer := internal.NewTokenizer(queue)
	tokenizer.Tokenize(text)
	tokenizer.Close()

	words := make([]string, 0, 10)
	for word := range queue {
		words = append(words, word)
	}
	return words
}

func TokenizeAndVerify(t *testing.T, text string, refWords ...string) {
	words := Tokenize(text)
	CompareStringSlices(t, words, refWords)
}
