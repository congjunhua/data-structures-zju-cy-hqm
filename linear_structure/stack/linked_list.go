package stack

// LinkedListStack 链表实现堆栈
type LinkedListStack struct {
	top *node
}

type node struct {
	value any
	next  *node
}

func (s *LinkedListStack) Empty() bool {
	return s.top == nil
}

func (s *LinkedListStack) Peek() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	return s.top.value, nil
}

func (s *LinkedListStack) Push(v any) {
	s.top = &node{value: v, next: s.top}
}

func (s *LinkedListStack) Pop() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	v := s.top.value
	s.top = s.top.next
	return v, nil
}
