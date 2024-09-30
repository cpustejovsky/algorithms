package linear_test

import (
	"testing"

	"github.com/cpustejovsky/algorithms/sorting/linear"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 42}
	searchNum := 42
	want := 7
	got := linear.Search(arr, searchNum)
	if want != *got {
		t.Fatalf("wanted %v; got %v", want, *got)
	}

}
