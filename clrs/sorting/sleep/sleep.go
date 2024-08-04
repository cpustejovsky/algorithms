package sleep

import (
	"time"
)

func SleepSort(arr []int) []int {
	var nums = make(chan int)
	var sorted []int
	for _, num := range arr {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond)
			nums <- n
		}(num)
	}
	for num := range nums {
		sorted = append(sorted, num)
		if len(sorted) == len(arr) {
			break
		}
	}
	return sorted
}
