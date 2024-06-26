package bst

type BinarySearchTree struct {
	Value  int
	Parent *BinarySearchTree
	Left   *BinarySearchTree
	Right  *BinarySearchTree
}

func Search(t *BinarySearchTree, value int) *BinarySearchTree {
	for t == nil && value != t.Value {
		if value < t.Value {
			t = t.Left
		} else {
			t = t.Right
		}
	}
	return t
}

func Walk(t *BinarySearchTree, arr *[]int) {
	if t != nil {
		Walk(t.Left, arr)
		*arr = append(*arr, t.Value)
		Walk(t.Right, arr)
	}
}

func Minimum(t *BinarySearchTree) *BinarySearchTree {
	for t.Left != nil {
		t = t.Left
	}
	return t
}

func Maximum(t *BinarySearchTree) *BinarySearchTree {
	for t.Right != nil {
		t = t.Right
	}
	return t
}

// Successor takes a BinarySearchTree and find the next node based on in-order traversal
func Successor(t *BinarySearchTree) *BinarySearchTree {
	if t.Right != nil {
		return Minimum(t.Right)
	}
	parent := t.Parent
	for parent != nil && t == parent.Right {
		t = parent
		parent = parent.Parent
	}
	return parent
}

// Predecessor takes a BinarySearchTree and find the preceding node based on in-order traversal
func Predecessor(t *BinarySearchTree) *BinarySearchTree {
	//TODO: test this to ensure it's correct as you cannot find a 1:1 implementation in CLRS or online
	if t.Left != nil {
		return Maximum(t.Left)
	}
	parent := t.Parent
	for parent != nil && t == parent.Left {
		t = parent
		parent = parent.Parent
	}
	return parent
}
