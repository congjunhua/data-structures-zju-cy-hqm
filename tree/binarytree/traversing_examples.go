package binarytree

import (
	"cmp"
	"fmt"
	"reflect"
)

// LeafNodes 给定一个二叉树，输出所有的叶子节点。
func (t *BinaryTree[T]) LeafNodes(vs *[]T) {
	t.leafNodes(t.Root, vs)
}

func (t *BinaryTree[T]) leafNodes(c *node[T], vs *[]T) {
	if c != nil {
		if c.left == nil && c.right == nil {
			*vs = append(*vs, c.value)
		}
		t.leafNodes(c.left, vs)
		t.leafNodes(c.right, vs)
	}
}

// Depth 计算给定二叉树的深度。
func (t *BinaryTree[T]) Depth() int {
	return t.Root.depth()
}

func (n *node[T]) depth() int {
	if n == nil {
		return 0
	}
	ld := n.left.depth()
	rd := n.right.depth()
	return max(ld, rd) + 1
}

/*
PrefixInfixPostfix
给定一个表示二元运算表达式的二叉树，其中叶子节点为运算数，非叶子节点为运算符。
依次返回「前缀表达式」、「中缀表达式」和「后缀表达式」。
*/
func (t *BinaryTree[T]) PrefixInfixPostfix() (string, string, string) {
	var vs []T
	t.TraverseInPreOrderByRecursion(&vs)
	prefix := ""
	for i := range vs {
		prefix += fmt.Sprintf("%v", vs[i])
	}

	infix := ""
	t.traverseInInOrderByRecursionWithBrackets(t.Root, &infix)

	vs = []T{}
	t.TraverseInPostOrderByRecursion(&vs)
	postfix := ""
	for i := range vs {
		postfix += fmt.Sprintf("%v", vs[i])
	}

	return prefix, infix, postfix
}

func (t *BinaryTree[T]) traverseInInOrderByRecursionWithBrackets(c *node[T], s *string) {
	if c != nil {
		if c.left != nil && c.right != nil {
			*s += "("
		}
		t.traverseInInOrderByRecursionWithBrackets(c.left, s)
		*s += fmt.Sprintf("%v", c.value)
		t.traverseInInOrderByRecursionWithBrackets(c.right, s)
		if c.left != nil && c.right != nil {
			*s += ")"
		}
	}
}

// RecoverFromPreorderAndInorder 根据前序和中序遍历的结果还原二叉树
func (t *BinaryTree[T]) RecoverFromPreorderAndInorder(preorder, inorder []T) {
	t.Root = recoverFromPreorderAndInorder(preorder, inorder)
}

func recoverFromPreorderAndInorder[T cmp.Ordered](preorder, inorder []T) *node[T] {
	if len(preorder) == 0 {
		return nil
	}

	rootValue := preorder[0]
	root := &node[T]{value: rootValue}

	// 在中序遍历结果中找到 rootValue 的下标
	index := 0
	for i, v := range inorder {
		if reflect.DeepEqual(v, rootValue) {
			index = i
			break
		}
	}

	// 递归地构建左子树和右子树
	root.left = recoverFromPreorderAndInorder(preorder[1:1+index], inorder[:index])
	root.right = recoverFromPreorderAndInorder(preorder[1+index:], inorder[index+1:])

	return root
}

// RecoverFromInorderAndPostorder 根据中序和后序遍历的结果还原二叉树
func (t *BinaryTree[T]) RecoverFromInorderAndPostorder(inorder, postorder []T) {
	t.Root = recoverFromInorderAndPostorder(inorder, postorder)
}

func recoverFromInorderAndPostorder[T cmp.Ordered](inorder, postorder []T) *node[T] {
	if len(postorder) == 0 {
		return nil
	}

	rootValue := postorder[len(postorder)-1]
	root := &node[T]{value: rootValue}

	// 在中序遍历结果中找到 rootValue 的下标
	index := 0
	for i, v := range inorder {
		if reflect.DeepEqual(v, rootValue) {
			index = i
			break
		}
	}

	// 递归地构建左子树和右子树
	root.left = recoverFromInorderAndPostorder(inorder[:index], postorder[:index])
	root.right = recoverFromInorderAndPostorder(
		inorder[index+1:], postorder[index:len(postorder)-1],
	)

	return root
}
