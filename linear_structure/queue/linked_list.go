package queue

// LinkedListQueue 链表实现队列
type LinkedListQueue struct {
	front, rear *node
}

type node struct {
	value any
	next  *node
}

func (l *LinkedListQueue) Empty() bool {
	return l.front == nil
}

func (l *LinkedListQueue) Peek() (any, error) {
	if l.Empty() {
		return nil, EmptyQueue
	}
	return l.front.value, nil
}

func (l *LinkedListQueue) Put(v any) {
	n := &node{
		value: v,
	}

	if l.Empty() {
		l.front = n
		l.rear = n
		return
	}

	l.rear.next = n
	l.rear = n
}

func (l *LinkedListQueue) Poll() (any, error) {
	if l.Empty() {
		return nil, EmptyQueue
	}

	v := l.front.value

	if l.front == l.rear {
		l.front = nil
		l.rear = nil
		return v, nil
	}

	nt := l.front.next
	l.front.next = nil
	l.front = nt

	return v, nil
}
