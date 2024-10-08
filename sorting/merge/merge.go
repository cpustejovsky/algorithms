package merge

import "github.com/cpustejovsky/algorithms/sorting/selection"

func Sort(nums []int, limit int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	if length < limit {
		return selection.Sort(nums)
	}
	middle := length / 2
	left := Sort(nums[:middle], limit)
	right := Sort(nums[middle:], limit)
	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	var i, j int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result
}
