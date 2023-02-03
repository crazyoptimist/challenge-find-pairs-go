package main

import (
	"testing"
)

func TestFindPairs(t *testing.T) {
	t.Run("it should find all the existing pairs", func(t *testing.T) {
		nums := []int{1, 9, 5, 0, 20, -4, 12, 16, 7}
		target := 12

		want := 3
		got := len(FindPairs(nums, target))

		if got != want {
			t.Errorf("expected %d pairs, but got %d pairs", want, got)
		}
	})

	t.Run("pairs should sum up to the given target", func(t *testing.T) {
		nums := []int{1, 9, 5, 0, 20, -4, 12, 16, 7}
		target := 12

		pairs := FindPairs(nums, target)

		for _, pair := range pairs {
			if pair[0]+pair[1] != target {
				t.Error("every pair should sum up to the target")
			}
		}
	})
}
