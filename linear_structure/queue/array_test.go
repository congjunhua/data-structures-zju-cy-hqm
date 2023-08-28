package queue

import (
	"errors"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue()

	if !q.Empty {
		t.Errorf("Newly created queue should be empty")
	}

	// Put elements into the queue
	for i := 0; i < capacity; i++ {
		if e := q.Put(i); e != nil {
			t.Errorf("Failed to put element %d into queue: %v", i, e)
		}
	}

	if !q.Full() {
		t.Errorf("Queue should be full after adding %v elements", capacity)
	}

	// Try to put more elements into the queue (should fail)
	if e := q.Put(10); e == nil || !errors.Is(e, FullQueue) {
		t.Errorf("Expected Put to fail with FullQueue error")
	}

	// Peek at the front element
	if f, e := q.Peek(); e != nil || f != 0 {
		t.Errorf("Peek did not return the expected result: %v, %v", f, e)
	}

	// Poll all elements from the queue
	for i := 0; i < capacity; i++ {
		if v, e := q.Poll(); e != nil || v != i {
			t.Errorf("Poll did not return the expected result: %v, %v", v, e)
		}
	}

	if !q.Empty {
		t.Errorf("Queue should be empty after polling all elements")
	}

	// Try to poll an element from an empty queue (should fail)
	if _, e := q.Poll(); e == nil || !errors.Is(e, EmptyQueue) {
		t.Errorf("Expected Poll to fail with EmptyQueue error")
	}
}
