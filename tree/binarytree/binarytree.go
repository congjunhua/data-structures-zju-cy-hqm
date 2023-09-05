package binarytree

import (
	"cmp"
)

// 链表实现二叉树

type BinaryTree[T cmp.Ordered] struct {
	Root *node[T]
}

type node[T cmp.Ordered] struct {
	value       T
	left, right *node[T]
}

func (t *BinaryTree[T]) Empty() bool {
	return t.Root == nil
}

func (t *BinaryTree[T]) Insert(v T) {
	if t.Root == nil {
		t.Root = &node[T]{value: v}
	} else {
		t.insert(t.Root, v)
	}
}

func (t *BinaryTree[T]) insert(n *node[T], v T) {
	if v < n.value {
		if n.left == nil {
			n.left = &node[T]{value: v}
		} else {
			t.insert(n.left, v)
		}
	} else {
		if n.right == nil {
			n.right = &node[T]{value: v}
		} else {
			t.insert(n.right, v)
		}
	}
}
