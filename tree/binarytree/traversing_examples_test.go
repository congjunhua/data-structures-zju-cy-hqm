package binarytree

import (
	"cmp"
	"reflect"
	"strings"
	"testing"
)

func newSkewedBinaryTree(depth int, left bool) BinaryTree[int] {
	bt := BinaryTree[int]{}
	if depth == 0 {
		return bt
	}
	n := &node[int]{value: 1}
	bt.Root = n
	if left {
		for i := 2; i <= depth; i++ {
			n.left = &node[int]{value: i}
			n = n.left
		}
	} else {
		for i := 2; i <= depth; i++ {
			n.right = &node[int]{value: i}
			n = n.right
		}
	}
	return bt
}

func newPerfectBinaryTree(depth int) BinaryTree[int] {
	bt := BinaryTree[int]{}
	if depth == 0 {
		return bt
	}
	root := &node[int]{value: 1}
	upstairs := []*node[int]{root}
	for i := 2; i <= depth; i++ {
		number := 1 << (i - 1) // 该层节点数
		nodes := make([]*node[int], number)
		for j := number; j < 1<<i; j++ {
			nodes[j-number] = &node[int]{value: j}
		}
		lrs := nodes
		for _, up := range upstairs {
			up.left = lrs[0]
			up.right = lrs[1]
			lrs = lrs[2:]
		}
		upstairs = nodes
	}
	bt.Root = root
	return bt
}

func TestBinaryTree_LeafNodes(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name     string
		tree     BinaryTree[T]
		expected []T
	}
	cases := []testCase[int]{
		{
			name:     "skewedBinaryTree",
			tree:     newSkewedBinaryTree(3, true),
			expected: []int{3},
		},
		{
			name: "completeBinaryTree",
			tree: func() BinaryTree[int] {
				root := &node[int]{value: 1}
				a := &node[int]{value: 2}
				b := &node[int]{value: 3}
				a.left = &node[int]{value: 4}
				a.right = &node[int]{value: 5}
				b.left = &node[int]{value: 6}
				root.left = a
				root.right = b
				bt := BinaryTree[int]{}
				bt.Root = root
				return bt
			}(),
			expected: []int{4, 5, 6},
		},
		{
			name:     "perfectBinaryTree",
			tree:     newPerfectBinaryTree(3),
			expected: []int{4, 5, 6, 7},
		},
	}
	for _, c := range cases {
		t.Run(
			c.name, func(t *testing.T) {
				var vs []int
				c.tree.LeafNodes(&vs)
				if !reflect.DeepEqual(vs, c.expected) {
					t.Errorf("%s Failed, Expected: %v, Got %v", c.name, c.expected, vs)
				}
			},
		)
	}
}

func TestBinaryTree_Depth(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name     string
		tree     BinaryTree[T]
		expected int
	}
	cases := []testCase[int]{
		{
			name:     "skewed_zero",
			tree:     newSkewedBinaryTree(0, true),
			expected: 0,
		},
		{
			name:     "skewed",
			tree:     newSkewedBinaryTree(10, false),
			expected: 10,
		},
		{
			name:     "perfect_zero",
			tree:     newPerfectBinaryTree(0),
			expected: 0,
		},
		{
			name:     "perfect",
			tree:     newPerfectBinaryTree(10),
			expected: 10,
		},
	}
	for _, c := range cases {
		t.Run(
			c.name, func(t *testing.T) {
				if got := c.tree.Depth(); got != c.expected {
					t.Errorf("Depth() = %v, expected %v", got, c.expected)
				}
			},
		)
	}
}

func TestBinaryTree_PrefixInfixPostfixNotation(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name            string
		tree            BinaryTree[T]
		expectedPrefix  string
		expectedInfix   string
		expectedPostfix string
	}
	cases := []testCase[string]{
		{
			name: "test",
			tree: func() BinaryTree[string] {
				root := &node[string]{value: "+"}
				root.left = &node[string]{value: "+"}
				root.left.left = &node[string]{value: "a"}
				root.left.right = &node[string]{value: "*"}
				root.left.right.left = &node[string]{value: "b"}
				root.left.right.right = &node[string]{value: "c"}
				root.right = &node[string]{value: "*"}
				root.right.left = &node[string]{value: "+"}
				root.right.left.left = &node[string]{value: "*"}
				root.right.left.left.left = &node[string]{value: "d"}
				root.right.left.left.right = &node[string]{value: "e"}
				root.right.left.right = &node[string]{value: "f"}
				root.right.right = &node[string]{value: "g"}
				return BinaryTree[string]{Root: root}
			}(),
			expectedPrefix:  "++a*bc*+*defg",
			expectedInfix:   "((a+(b*c))+(((d*e)+f)*g))",
			expectedPostfix: "abc*+de*f+g*+",
		},
	}
	for _, c := range cases {
		t.Run(
			c.name, func(t *testing.T) {
				prefix, infix, postfix := c.tree.PrefixInfixPostfix()
				if prefix != c.expectedPrefix {
					t.Errorf("【prefix】expected: %v, got: %v", c.expectedPrefix, prefix)
				}
				if infix != c.expectedInfix {
					t.Errorf("【infix】expected: %v, got: %v", c.expectedInfix, infix)
				}
				if postfix != c.expectedPostfix {
					t.Errorf("【postfix】expected: %v, got: %v", c.expectedPostfix, postfix)
				}
			},
		)
	}
}

func TestBinaryTree_RecoverFromPreorderAndInorder(t *testing.T) {
	type cases[T cmp.Ordered] struct {
		name     string
		tree     BinaryTree[T]
		preorder []T
		inorder  []T
	}
	tests := []cases[string]{
		{
			name:     "test",
			preorder: strings.Split("abcdefghij", ""),
			inorder:  strings.Split("cbedahgijf", ""),
		},
	}
	for _, c := range tests {
		t.Run(
			c.name, func(t *testing.T) {
				c.tree.RecoverFromPreorderAndInorder(c.preorder, c.inorder)
				var pre, in []string
				c.tree.TraverseInPreOrderByRecursion(&pre)
				if !reflect.DeepEqual(c.preorder, pre) {
					t.Errorf("Failed Expected %v, Got %v\n", c.preorder, pre)
				}
				c.tree.TraverseInInOrderByRecursion(&in)
				if !reflect.DeepEqual(c.inorder, in) {
					t.Errorf("Failed Expected %v, Got %v\n", c.inorder, in)
				}
			},
		)
	}
}

func TestBinaryTree_RecoverFromInorderAndPostorder(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name      string
		tree      BinaryTree[T]
		inorder   []T
		postorder []T
	}
	cases := []testCase[string]{
		{
			name:      "test",
			inorder:   strings.Split("cbedahgijf", ""),
			postorder: strings.Split("cedbhjigfa", ""),
		},
	}
	for _, c := range cases {
		t.Run(
			c.name, func(t1 *testing.T) {
				c.tree.RecoverFromInorderAndPostorder(c.inorder, c.postorder)
				var in, post []string
				c.tree.TraverseInInOrderByRecursion(&in)
				if !reflect.DeepEqual(c.inorder, in) {
					t.Errorf("Failed Expected %v, Got %v\n", c.inorder, in)
				}
				c.tree.TraverseInPostOrderByRecursion(&post)
				if !reflect.DeepEqual(c.postorder, post) {
					t.Errorf("Failed Expected %v, Got %v\n", c.postorder, post)
				}
			},
		)
	}
}
