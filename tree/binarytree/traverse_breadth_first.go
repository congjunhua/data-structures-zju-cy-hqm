package binarytree

import (
	"data-structures-zju-cy-hqm/linear_structure/queue"
	"data-structures-zju-cy-hqm/linear_structure/stack"
)

// 广度优先遍历

// TraverseBreadthFirstByQueue 广度优先遍历（基于队列）
func (t *BinaryTree[T]) TraverseBreadthFirstByQueue(vs *[]T) error {
	q := &queue.LinkedListQueue{}
	q.Put(t.Root)
	for !q.Empty() {
		a, e := q.Poll()
		if e != nil {
			return e
		}
		n := a.(*node[T])
		*vs = append(*vs, n.value)
		if n.left != nil {
			q.Put(n.left)
		}
		if n.right != nil {
			q.Put(n.right)
		}
	}
	return nil
}

// TraverseBreadthFirstByStack 广度优先遍历（基于堆栈）
func (t *BinaryTree[T]) TraverseBreadthFirstByStack(vs *[]T) error {
	current, next := &stack.LinkedListStack{}, &stack.LinkedListStack{}
	current.Push(t.Root)
	n := new(node[T])
	for !current.Empty() {
		if v, e := current.Pop(); e != nil {
			return e
		} else {
			n = v.(*node[T])
		}
		*vs = append(*vs, n.value)
		if n.right != nil {
			next.Push(n.right)
		}
		if n.left != nil {
			next.Push(n.left)
		}
		if current.Empty() {
			current, next = next, current
		}
	}
	return nil
}
