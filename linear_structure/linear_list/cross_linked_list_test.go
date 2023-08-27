package linear_list

import (
	"testing"
)

func TestCrossLinkedList(t *testing.T) {
	l := NewCrossLinkedList(3, 5)

	// Test Set and Get
	err := l.Set(1, 1, 1)
	if err != nil {
		t.Errorf("Error setting element: %v", err)
	}
	elem, e := l.Get(1, 1)
	if e != nil {
		t.Errorf("Error getting element: %v", e)
	}
	if elem.Value != 1 {
		t.Errorf("Expected value 1, got %v", elem.Value)
	}

	// Test Delete
	err = l.Delete(1, 1)
	if err != nil {
		t.Errorf("Error deleting element: %v", err)
	}
	_, err = l.Get(1, 1)
	if err == nil {
		t.Errorf("Element should have been deleted")
	}
}

func TestOutOfRange(t *testing.T) {
	l := NewCrossLinkedList(3, 5)

	if !l.outOfRange(0, 1) {
		t.Errorf("Expected outOfRange to return true for row 0")
	}
	if !l.outOfRange(4, 1) {
		t.Errorf("Expected outOfRange to return true for row 4")
	}
	if !l.outOfRange(1, 0) {
		t.Errorf("Expected outOfRange to return true for column 0")
	}
	if !l.outOfRange(1, 6) {
		t.Errorf("Expected outOfRange to return true for column 6")
	}
	if l.outOfRange(1, 1) {
		t.Errorf("Expected outOfRange to return false for valid row and column")
	}
}
