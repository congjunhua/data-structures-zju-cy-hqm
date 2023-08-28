package queue

import (
	"errors"
	"testing"
)

func TestLinkedListQueue(t *testing.T) {
	q := &LinkedListQueue{}

	if !q.Empty() {
		t.Errorf("Newly created queue should be empty")
	}

	// Put elements into the queue
	for i := 1; i <= 5; i++ {
		q.Put(i)
	}

	// Peek at the front element
	if v, e := q.Peek(); e != nil || v != 1 {
		t.Errorf("Peek did not return the expected result: %v, %v", v, e)
	}

	// Poll elements from the queue
	for i := 1; i <= 5; i++ {
		if v, e := q.Poll(); e != nil || v != i {
			t.Errorf("Poll did not return the expected result: %v, %v", v, e)
		}
	}

	if !q.Empty() {
		t.Errorf("Queue should be empty after polling all elements")
	}

	// Try to poll an element from an empty queue (should fail)
	if _, e := q.Poll(); e == nil || !errors.Is(e, EmptyQueue) {
		t.Errorf("Expected Poll to fail with EmptyQueue error")
	}
}

func TestLinkedListQueueMultiplePuts(t *testing.T) {
	q := &LinkedListQueue{}

	// Put multiple elements into the queue
	for i := 1; i <= 10; i++ {
		q.Put(i)
	}

	// Poll elements from the queue
	for i := 1; i <= 10; i++ {
		if v, e := q.Poll(); e != nil || v != i {
			t.Errorf("Poll did not return the expected result: %v, %v", v, e)
		}
	}
}
