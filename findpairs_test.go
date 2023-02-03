package main

import (
	"reflect"
	"testing"
)

func TestFindPairs(t *testing.T) {
	sampleSlice := []int{1, 9, 5, 0, 20, -4, 12, 16, 7}
	sampleSum := 12

	got := FindPairs(sampleSlice, sampleSum)
	want := [][]int{{12, 0}, {5, 7}, {16, -4}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}
