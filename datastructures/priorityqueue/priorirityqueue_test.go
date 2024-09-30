package priorityqueue_test

import (
	"container/heap"
	"math/rand"
	"sort"
	"testing"

	"github.com/cpustejovsky/algorithms/datastructures/priorityqueue"
)

func TestPriorityQueue(t *testing.T) {
	var p priorityqueue.PriorityQueue
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = rand.Intn(1000)
	}
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	for i := 0; i < len(nums); i++ {
		p.Insert(nums[i])
	}
	i := len(sorted) - 1
	for p.Len() > 0 {
		x := p.Pop()
		if x != sorted[i] {
			t.Fatalf("expect %d, got %d", sorted[i], x)
		}
		i--
	}
}

func TestPriorityQueueHeap(t *testing.T) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = rand.Intn(1000)
	}
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for _, num := range nums {
		heap.Push(maxHeap, num)
	}
	i := len(sorted) - 1
	for maxHeap.Len() > 0 {
		x := heap.Pop(maxHeap).(int)
		if x != sorted[i] {
			t.Fatalf("expect %d, got %d", sorted[i], x)
		}
		i--
	}

}

func BenchmarkPriorityQueue(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = rand.Intn(1000)
	}
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	for i := 0; i < b.N; i++ {
		var p priorityqueue.PriorityQueue
		for _, num := range nums {
			p.Insert(num)
		}
		i := len(sorted) - 1
		for p.Len() > 0 {
			x := p.Pop()
			if x != sorted[i] {
				b.Fatalf("expect %d, got %d", sorted[i], x)
			}
			i--
		}
	}
}

// A MaxHeap to hold the profits (we implement this as a min-heap with inverted profits)
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Reverses the order to create a max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BenchmarkPriorityQueueHeap(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = rand.Intn(1000)
	}
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	for i := 0; i < b.N; i++ {
		maxHeap := &MaxHeap{}
		heap.Init(maxHeap)
		for _, num := range nums {
			heap.Push(maxHeap, num)
		}
		i := len(sorted) - 1
		for maxHeap.Len() > 0 {
			x := heap.Pop(maxHeap).(int)
			if x != sorted[i] {
				b.Fatalf("expect %d, got %d", sorted[i], x)
			}
			i--
		}
	}

}
