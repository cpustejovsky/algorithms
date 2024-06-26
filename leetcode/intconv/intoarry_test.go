package intconv_test

import (
	"github.com/cpustejovsky/algorithms/leetcode/intconv"
	"testing"
)

func TestIntToSlice(t *testing.T) {
	t.Run("positive integer", func(t *testing.T) {
		num := 123
		want := []int{1, 2, 3}
		got := intconv.IntToSlice(num)
		if !EqualSlices(got, want) {
			t.Fatalf("wanted %v; got %v", want, got)
		}
	})
	t.Run("negative integer", func(t *testing.T) {
		num := -123
		want := []int{-1, -2, -3}
		got := intconv.IntToSlice(num)
		if !EqualSlices(got, want) {
			t.Fatalf("wanted %v; got %v", want, got)
		}
	})
}

func TestSliceToInt(t *testing.T) {
	t.Run("with positive numbers in int slice", func(t *testing.T) {
		nums := []int{1, 2, 3}
		want := 123
		got := intconv.SliceToInt(nums)
		if got != want {
			t.Fatalf("wanted %v; got %v", want, got)
		}
	})
	t.Run("with negative numbers in int slice", func(t *testing.T) {
		nums := []int{-1, -2, -3}
		want := -123
		got := intconv.SliceToInt(nums)
		if got != want {
			t.Fatalf("wanted %v; got %v", want, got)
		}
	})
}

func EqualSlices[C comparable](x, y []C) bool {
	if len(x) != len(y) {
		return false
	}
	for i, val := range x {
		if y[i] != val {
			return false
		}
	}
	return true
}
