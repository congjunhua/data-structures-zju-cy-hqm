package queue

import (
	"errors"
)

// 使用数组实现循环队列

const capacity = 4

var (
	EmptyQueue = errors.New("对列为空")
	FullQueue  = errors.New("对列已满")
)

type ArrayQueue struct {
	Values      [capacity]any
	Empty       bool
	Front, Rear uint
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		Empty: true,
	}
}

func (q *ArrayQueue) Full() bool {
	return next(q.Rear) == q.Front
}

func (q *ArrayQueue) Peek() (any, error) {
	if q.Empty {
		return nil, EmptyQueue
	}
	return q.Values[q.Front], nil
}

func (q *ArrayQueue) Put(v any) error {
	if q.Full() {
		return FullQueue
	}
	if !q.Empty {
		q.Rear++
	}
	q.Values[q.Rear] = v
	if q.Empty {
		q.Empty = !q.Empty
	}
	return nil
}

func (q *ArrayQueue) Poll() (any, error) {
	if q.Empty {
		return nil, EmptyQueue
	}
	v := q.Values[q.Front]
	q.Values[q.Front] = nil
	if q.Front != q.Rear {
		q.Front = next(q.Front)
	} else {
		q.Empty = !q.Empty
	}
	return v, nil
}

func next(p uint) uint {
	return (p + 1) % capacity
}
