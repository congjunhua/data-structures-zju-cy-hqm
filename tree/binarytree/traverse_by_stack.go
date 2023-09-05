package binarytree

import "data-structures-zju-cy-hqm/linear_structure/stack"

// 基于堆栈遍历二叉树

// TraverseInInOrderByStack 中序遍历
func (t *BinaryTree[T]) TraverseInInOrderByStack(vs *[]T) error {
	n, s := t.Root, &stack.LinkedListStack{}
	for {
		if n.left != nil {
			s.Push(n)
			n = n.left
			continue
		}
		*vs = append(*vs, n.value)
		if n.right != nil {
			n = n.right
			continue
		}
		if s.Empty() {
			break
		}
		p, e := s.Pop()
		if e != nil {
			return e
		}
		n = &node[T]{value: p.(*node[T]).value, right: p.(*node[T]).right}
	}
	return nil
}

// TraverseInPostOrderByStack 后序遍历
func (t *BinaryTree[T]) TraverseInPostOrderByStack(vs *[]T) error {
	n, s := t.Root, &stack.LinkedListStack{}
	for {
		if n.left != nil {
			s.Push(&node[T]{value: n.value, right: n.right})
			n = n.left
			continue
		}
		if n.right != nil {
			r := n.right
			n.right = nil
			s.Push(n)
			n = r
			continue
		}
		*vs = append(*vs, n.value)
		if s.Empty() {
			break
		}
		p, e := s.Pop()
		if e != nil {
			return e
		}
		n = p.(*node[T])
	}
	return nil
}
