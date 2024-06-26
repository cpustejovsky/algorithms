package intconv

func IntToSlice(x int) []int {
	var res []int
	for x != 0 {
		y := x % 10
		x /= 10
		res = append([]int{y}, res...)
	}
	return res
}

func SliceToInt(nums []int) int {
	var total int
	pow := len(nums) - 1
	for i, num := range nums {
		total += num * pow10(pow-i)
	}
	return total
}

func pow10(pow int) int {
	x := 1
	for i := 0; i < pow; i++ {
		x *= 10
	}
	return x
}
