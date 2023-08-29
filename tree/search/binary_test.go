package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}

	// Test cases for elements present in the array
	if idx := BinarySearch(arr, 5); idx != 2 {
		t.Errorf("Expected index 2 for element 5, but got %d", idx)
	}
	if idx := BinarySearch(arr, 15); idx != 7 {
		t.Errorf("Expected index 7 for element 15, but got %d", idx)
	}

	// Test cases for elements not present in the array
	if idx := BinarySearch(arr, 2); idx != -1 {
		t.Errorf("Expected index -1 for element 2, but got %d", idx)
	}
	if idx := BinarySearch(arr, 6); idx != -1 {
		t.Errorf("Expected index -1 for element 6, but got %d", idx)
	}
}
