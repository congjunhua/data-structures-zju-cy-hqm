package binarytree

import (
	"reflect"
	"testing"
)

func TestTraverseByStack(t *testing.T) {
	// 创建节点实例
	node1 := &node[int]{value: 1}
	node2 := &node[int]{value: 2}
	node3 := &node[int]{value: 3}
	node4 := &node[int]{value: 4}
	node5 := &node[int]{value: 5}

	// 链接节点形成二叉树结构
	node1.left = node2
	node1.right = node3
	node2.left = node4
	node2.right = node5

	// 创建二叉树实例
	tree := BinaryTree[int]{Root: node1}

	inOrder(t, tree, []int{4, 2, 5, 1, 3})
	postOrder(t, tree, []int{4, 5, 2, 3, 1})

	node5.left = &node[int]{value: 6}
	inOrder(t, tree, []int{4, 2, 6, 5, 1, 3})
	postOrder(t, tree, []int{4, 6, 5, 2, 3, 1})

	node5.left = nil
	node5.right = &node[int]{value: 6}
	inOrder(t, tree, []int{4, 2, 5, 6, 1, 3})
	postOrder(t, tree, []int{4, 6, 5, 2, 3, 1})

	node5.left = &node[int]{value: 6}
	node5.right = &node[int]{value: 7}
	inOrder(t, tree, []int{4, 2, 6, 5, 7, 1, 3})
	postOrder(t, tree, []int{4, 6, 7, 5, 2, 3, 1})

	node2.right = nil
	inOrder(t, tree, []int{4, 2, 1, 3})
	postOrder(t, tree, []int{4, 2, 3, 1})

	node1.right = nil
	inOrder(t, tree, []int{4, 2, 1})
	postOrder(t, tree, []int{4, 2, 1})

	node1.left = nil
	inOrder(t, tree, []int{1})
	postOrder(t, tree, []int{1})
}

func inOrder(t *testing.T, tree BinaryTree[int], expected []int) {
	result := make([]int, 0)
	if e := tree.TraverseInInOrderByStack(&result); e != nil {
		t.Error("遍历出错：", e)
		return
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed. Expected: %v, Got: %v", expected, result)
	}
}

func postOrder(t *testing.T, tree BinaryTree[int], expected []int) {
	result := make([]int, 0)
	if e := tree.TraverseInPostOrderByStack(&result); e != nil {
		t.Error("遍历出错：", e)
		return
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed. Expected: %v, Got: %v", expected, result)
	}
}
