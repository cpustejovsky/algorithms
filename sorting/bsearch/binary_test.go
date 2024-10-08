package bsearch_test

import (
	"testing"

	"github.com/cpustejovsky/algorithms/sorting/bsearch"
)

func TestSearch(t *testing.T) {
	var arr = []int{1, 2, 3, 5, 7, 9, 11, 15, 16}
	t.Run("Match", func(t *testing.T) {
		want := 4
		got := bsearch.Search(7, arr)
		if got != want {
			t.Fatalf("got:\t%v\nwanted:\t%v\n", got, want)
		}
	})
	t.Run("No Match", func(t *testing.T) {
		want := 9
		got := bsearch.Search(42, arr)
		if got != want {
			t.Fatalf("got:\t%v\nwanted:\t%v\n", got, want)
		}
	})
}

func BenchmarkSearch(b *testing.B) {
	var arr = []int{1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 7, 9, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 15}
	for i := 0; i < b.N; i++ {
		want := 52
		got := bsearch.Search(7, arr)
		if got != want {
			b.Fatalf("got:\t%v\nwanted:\t%v\n", got, want)
		}
	}
}
