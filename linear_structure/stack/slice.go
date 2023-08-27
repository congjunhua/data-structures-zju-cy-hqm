package stack

// 切片实现堆栈

type SliceStack struct {
	Values []any
}

func NewSliceStack() *SliceStack {
	return &SliceStack{}
}

func (s *SliceStack) Empty() bool {
	return len(s.Values) == 0
}

// 因为切片是动态数组，可以不关心栈满，故不实现判满。
// func (s *SliceStack) Full() bool {}

func (s *SliceStack) Peek() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	return s.Values[len(s.Values)-1], nil
}

func (s *SliceStack) Push(v any) {
	s.Values = append(s.Values, v)
}

func (s *SliceStack) Pop() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	i := len(s.Values) - 1
	v := s.Values[i]
	s.Values = s.Values[:i]
	return v, nil
}
