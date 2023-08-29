package search

import (
	"testing"
)

func TestSequentialSearch(t *testing.T) {
	arr := []any{1, 3, 5, 7, 9}

	// Test cases for elements present in the array
	if idx := SequentialSearch(arr, 5); idx != 2 {
		t.Errorf("Expected index 2 for element 5, but got %d", idx)
	}
	if idx := SequentialSearch(arr, 7); idx != 3 {
		t.Errorf("Expected index 3 for element 7, but got %d", idx)
	}

	// Test cases for elements not present in the array
	if idx := SequentialSearch(arr, 2); idx != -1 {
		t.Errorf("Expected index -1 for element 2, but got %d", idx)
	}
	if idx := SequentialSearch(arr, 6); idx != -1 {
		t.Errorf("Expected index -1 for element 6, but got %d", idx)
	}
}
