package binarytree

// 基于递归遍历二叉树

// TraverseInPreOrderByRecursion 前序遍历
func (t *BinaryTree[T]) TraverseInPreOrderByRecursion(vs *[]T) {
	t.traverseInPreOrderByRecursion(t.Root, vs)
}

func (t *BinaryTree[T]) traverseInPreOrderByRecursion(c *node[T], vs *[]T) {
	if c != nil {
		*vs = append(*vs, c.value)
		t.traverseInPreOrderByRecursion(c.left, vs)
		t.traverseInPreOrderByRecursion(c.right, vs)
	}
}

// TraverseInInOrderByRecursion 中序遍历
func (t *BinaryTree[T]) TraverseInInOrderByRecursion(vs *[]T) {
	t.traverseInInOrderByRecursion(t.Root, vs)
}

func (t *BinaryTree[T]) traverseInInOrderByRecursion(c *node[T], vs *[]T) {
	if c != nil {
		t.traverseInInOrderByRecursion(c.left, vs)
		*vs = append(*vs, c.value)
		t.traverseInInOrderByRecursion(c.right, vs)
	}
}

// TraverseInPostOrderByRecursion 后序遍历
func (t *BinaryTree[T]) TraverseInPostOrderByRecursion(vs *[]T) {
	t.traverseInPostOrderByRecursion(t.Root, vs)
}

func (t *BinaryTree[T]) traverseInPostOrderByRecursion(c *node[T], vs *[]T) {
	if c != nil {
		t.traverseInPostOrderByRecursion(c.left, vs)
		t.traverseInPostOrderByRecursion(c.right, vs)
		*vs = append(*vs, c.value)
	}
}
