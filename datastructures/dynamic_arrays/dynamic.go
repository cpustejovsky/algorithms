package dynamic_arrays

type DynamicArray struct {
	array []int
	index int
	size  int
}

func New(size int) *DynamicArray {
	a := make([]int, size)
	return &DynamicArray{
		array: a,
		index: -1,
		size:  size,
	}
}

func (d *DynamicArray) Len() int {
	return len(d.array)
}

func (d *DynamicArray) Cap() int {
	return cap(d.array)
}
func (d *DynamicArray) Index(i int) int {
	if i > d.index {
		return -1
	}
	return d.array[i]
}

func (d *DynamicArray) Insert(x int) {
	//check if the next index number is greater than the total capacity
	//an index 10 would require a capacity of 11 since the index starts at 0
	if d.index+1 > d.Cap()-1 {
		//Create new array with 2x capacity
		n := make([]int, d.size*2)
		//Copy old array into new array
		copy(n, d.array)
		d.array = n
		d.size = cap(n)
	}
	d.index++
	d.array[d.index] = x
}
