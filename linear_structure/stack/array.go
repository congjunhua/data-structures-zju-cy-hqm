package stack

import "errors"

// 数组实现堆栈

const capacity = 10

var (
	FullError  = errors.New("栈满")
	EmptyError = errors.New("栈空")
)

type ArrayStack struct {
	Values [capacity]any
	Top    int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		Top: -1,
	}
}

func (s *ArrayStack) Empty() bool {
	return s.Top == -1
}

func (s *ArrayStack) Full() bool {
	return s.Top == len(s.Values)-1
}

func (s *ArrayStack) Peek() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	return s.Values[s.Top], nil
}

func (s *ArrayStack) Push(v any) error {
	if s.Full() {
		return FullError
	}
	s.Top++
	s.Values[s.Top] = v
	return nil
}

func (s *ArrayStack) Pop() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	v := s.Values[s.Top]
	s.Top--
	return v, nil
}
