package priorityqueue

import "sort"

type PriorityQueue struct {
	values []int
}

func (p *PriorityQueue) Insert(val int) {
	p.values = append(p.values, val)
	sort.Ints(p.values)
}

func (p *PriorityQueue) Len() int {
	return len(p.values)
}

func (p *PriorityQueue) Pop() int {
	x := p.values[len(p.values)-1]
	p.values = p.values[:len(p.values)-1]
	return x
}
