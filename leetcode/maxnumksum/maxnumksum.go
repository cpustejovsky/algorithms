package maxnumksum

import "sort"

func MaxOperations(arr []int, sum int) int {
	sort.Ints(arr)
	var count, left, right int
	right = len(arr) - 1
	for left < right {
		candidate := arr[left] + arr[right]
		if candidate < sum {
			left++
		} else if candidate > sum {
			right--
		} else if candidate == sum {
			left++
			right--
			count++
		}
	}
	return count
}

func MaxOperationsMap(arr []int, sum int) int {
	countMap := make(map[int]int)
	var count int
	for _, current := range arr {
		complement := sum - current
		complementCount, _ := countMap[complement]
		if complementCount > 0 {
			//remove complement from map
			countMap[complement]--
			count++
		} else {
			countMap[current]++
		}
	}
	return count
}

// Time Limit Exceeded 47 / 51 testcases passed
// Attempted brute force solution of O(n^2) time complexity and met this limitation
// Checked on local and does pass but is ~4000 times slower than two pointer solution
func MaxOperationsSlow(nums []int, k int) int {
	var pairs int
	var pairNeeded int
	var i, j int
	for i < len(nums) {
		pairNeeded = k - nums[i]
		var pairFound bool
		j = i + 1
		for j < len(nums) {
			if nums[j] == pairNeeded {
				pairFound = true
				pairs++
				nums = removeAtIndex(nums, j)
				nums = removeAtIndex(nums, i)
				break
			} else {
				j++
			}
		}
		if !pairFound {
			i++
		}
	}
	return pairs
}

func removeAtIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func numSteps(s string) int {
	binary := []byte(s)
	var steps int
	lastIndex := len(s) - 1

	for len(binary[0:lastIndex+1]) > 1 {
		steps++
		offset := 1
		if binary[lastIndex] == '1' {
			binary[lastIndex] = '0'
			//if the next index is 0, change it to 1 and stop
			//else change the 1 to a 0 and increment index
			for {
				if offset == lastIndex {
					binary = append([]byte{'0'}, binary...)
					break
				} else if binary[lastIndex-offset] == '0' {
					binary[lastIndex-offset] = '1'
					break
				} else {
					binary[lastIndex-offset] = '0'
					offset++
				}
			}
		} else {
			lastIndex--
		}
	}
	return steps
}
