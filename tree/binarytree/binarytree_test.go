package binarytree

import (
	"reflect"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	bt := &BinaryTree[int]{}

	if !bt.Empty() {
		t.Error("Expected tree to be empty")
	}

	elements := []int{5, 3, 8, 2, 4, 7, 9, 1, 6}

	for _, v := range elements {
		bt.Insert(v)
	}

	if bt.Empty() {
		t.Error("Expected tree to be non-empty")
	}

	var preOrder, inOrder, postOrder []int
	expectedPreOrder := []int{5, 3, 2, 1, 4, 8, 7, 6, 9}
	expectedInOrder := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedPostOrder := []int{1, 2, 4, 3, 6, 7, 9, 8, 5}

	bt.TraverseInPreOrderByRecursion(&preOrder)
	bt.TraverseInInOrderByRecursion(&inOrder)
	bt.TraverseInPostOrderByRecursion(&postOrder)

	if !reflect.DeepEqual(preOrder, expectedPreOrder) {
		t.Error("Pre-order traversal mismatch")
	}

	if !reflect.DeepEqual(inOrder, expectedInOrder) {
		t.Error("In-order traversal mismatch")
	}

	if !reflect.DeepEqual(postOrder, expectedPostOrder) {
		t.Error("Post-order traversal mismatch")
	}
}
