package binarytree

import (
	"cmp"
	"reflect"
	"testing"
)

func TestBinaryTree_TraverseInBreadthFirstOrderByQueue(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name     string
		tree     BinaryTree[T]
		expected []T
	}
	cases := []testCase[int]{
		{
			name: "perfect",
			tree: func() BinaryTree[int] {
				root := &node[int]{value: 1}
				a := &node[int]{value: 2}
				b := &node[int]{value: 3}
				a.left = &node[int]{value: 4}
				a.right = &node[int]{value: 5}
				b.left = &node[int]{value: 6}
				b.right = &node[int]{value: 7}
				root.left = a
				root.right = b
				bt := BinaryTree[int]{}
				bt.Root = root
				return bt
			}(),
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, c := range cases {
		t.Run(
			c.name, func(t *testing.T) {
				var vs []int
				if e := c.tree.TraverseBreadthFirstByQueue(&vs); e != nil {
					t.Error(e)
				}
				if !reflect.DeepEqual(vs, c.expected) {
					t.Errorf("%s Failed, Expected: %v, Got %v", c.name, c.expected, vs)
				}
			},
		)
	}
}

func TestBinaryTree_TraverseInBreadthFirstOrderByStack(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name     string
		tree     BinaryTree[T]
		expected []T
	}
	cases := []testCase[int]{
		{
			name: "perfect",
			tree: func() BinaryTree[int] {
				root := &node[int]{value: 1}
				a := &node[int]{value: 2}
				b := &node[int]{value: 3}
				a.left = &node[int]{value: 4}
				a.right = &node[int]{value: 5}
				b.left = &node[int]{value: 6}
				b.right = &node[int]{value: 7}
				root.left = a
				root.right = b
				bt := BinaryTree[int]{}
				bt.Root = root
				return bt
			}(),
			expected: []int{1, 2, 3, 6, 7, 4, 5},
		},
	}
	for _, c := range cases {
		t.Run(
			c.name, func(t *testing.T) {
				var vs []int
				if e := c.tree.TraverseBreadthFirstByStack(&vs); e != nil {
					t.Error(e)
				}
				if !reflect.DeepEqual(vs, c.expected) {
					t.Errorf("%s Failed, Expected: %v, Got %v", c.name, c.expected, vs)
				}
			},
		)
	}
}
