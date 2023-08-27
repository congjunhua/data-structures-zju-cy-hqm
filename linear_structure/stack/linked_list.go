package stack

// 链表实现堆栈

type Node struct {
	Value any
	Next  *Node
}

type LinkedListStack struct {
	Top *Node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{}
}

func (s *LinkedListStack) Empty() bool {
	return s.Top == nil
}

func (s *LinkedListStack) Peek() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	return s.Top.Value, nil
}

func (s *LinkedListStack) Push(v any) {
	s.Top = &Node{Value: v, Next: s.Top}
}

func (s *LinkedListStack) Pop() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	v := s.Top.Value
	s.Top = s.Top.Next
	return v, nil
}
