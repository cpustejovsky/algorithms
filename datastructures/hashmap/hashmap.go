package hashmap

import (
	"hash/fnv"

	"github.com/cpustejovsky/algorithms/datastructures/linkedlist"
)

type HashMap[T comparable] []linkedlist.Node[T]

func New[T comparable](size int) HashMap[T] {
	return make([]linkedlist.Node[T], size)
}

func hashFunction(key string, size int) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	return int(hasher.Sum32()) % size
}

func (h *HashMap[T]) Insert(key string, value string) {

}
