package set

type Set interface {
	// Search returns the item with matching key in Set
	Search(any) any
	// Insert takes an iterable value and adds it to Set
	Insert(any)
	// Remove and return the store item with matching key
	Delete(any) any
	Minimum() any
	// Return stored item with smallest key
	Maximum() any
	// Return store itme with smallest key largest than key provided
	Successor(any) any
	// Return store itme with smallest key smallest than key provided
	Predecessor(any) any
	// Return sorted items in key order
	Iter_Order() []any
	// Return stored item with largest key
	// Build takes any sequence and creates a Set
	Build([]any)
	// Len returns the length of the Set
	Len() int
}

/*
An unsorted dynamic array would take O(n) time for all the operations
A sorted array would take the following
	build takes O(n log n)
	insert and delete take O(n)
	find takes O(log n)
	find_prev and find_next take O(log n)
	find_mind and find_max take O(1)
*/
