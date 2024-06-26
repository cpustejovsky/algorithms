package Ch2

import "github.com/cpustejovsky/algorithms/clrs"

func Sum(nums []int, sum int) bool {
	//Use Merge Sort that that's O(n lg n)
	sorted := clrs.MergeSort(nums, 4)
	//Set values for index
	i := 0
	j := len(sorted) - 1
	for i < j {
		if sorted[i]+sorted[j] == sum {
			return true
		} else if sorted[i]+sorted[j] > sum {
			j--
		} else {
			i++
		}
	}
	return false
}
