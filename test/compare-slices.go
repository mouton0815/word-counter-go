package test

import (
	"log"
	"reflect"
	"testing"
	"word-counter-go/main/internal"
)

func CompareStringSlices(t *testing.T, slice, refSlice []string) {
	if !reflect.DeepEqual(slice, EnsureNonNilStringSlice(refSlice)) {
		log.Printf("'%v' <-> '%v'\n", slice, refSlice)
		t.Fail()
	}
}

func CompareWordCountSlices(t *testing.T, slice, refSlice internal.WordCountSlice) {
	if !reflect.DeepEqual(slice, refSlice) {
		log.Printf("'%v' <-> '%v'\n", slice, refSlice)
		t.Fail()
	}
}

func EnsureNonNilStringSlice(slice []string) []string {
	if slice == nil {
		return []string{}
	}
	return slice
}
