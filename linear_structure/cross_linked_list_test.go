package linear_structure

import (
	"testing"
)

func TestCrossLinkedList(t *testing.T) {
	l := newCrossLinkedList(3, 5)

	// Test set and get
	err := l.set(1, 1, 1)
	if err != nil {
		t.Errorf("Error setting element: %v", err)
	}
	elem, e := l.get(1, 1)
	if e != nil {
		t.Errorf("Error getting element: %v", e)
	}
	if elem.value != 1 {
		t.Errorf("Expected value 1, got %v", elem.value)
	}

	// Test delete
	err = l.delete(1, 1)
	if err != nil {
		t.Errorf("Error deleting element: %v", err)
	}
	_, err = l.get(1, 1)
	if err == nil {
		t.Errorf("Element should have been deleted")
	}
}

func TestOutOfRange(t *testing.T) {
	l := newCrossLinkedList(3, 5)

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
