package binarytree

// 链表实现二叉树

type node struct {
	value       any
	left, right *node
}

type BinaryTree node
