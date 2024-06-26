package bitflip

func BruteForce(nums []int, k int) int {
	//sliding window of size k
	var left, right, count int
	right += k
	for right <= len(nums) {
		if nums[left] == 0 {
			for i := 0; i < len(nums[left:right]); i++ {
				if nums[left+i] == 0 {
					nums[left+i] = 1
				} else {
					nums[left+i] = 0
				}
			}
			count++
		}
		left++
		right++
	}
	for _, bit := range nums {
		if bit == 0 {
			return -1
		}
	}
	return count
}

func ConstantSpace(nums []int, k int) int {
	n := len(nums)
	if k > n {
		return -1
	}
	count := 0
	flip := 0

	for i := 0; i < n; i++ {
		if i >= k {
			flip ^= 1 - nums[i-k]
		}
		if nums[i] == flip {
			if i+k > n {
				return -1
			}
			nums[i] = 0
			flip ^= 1
			count++

		} else {
			nums[i] = 1
		}
	}

	return count
}
