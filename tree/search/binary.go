package search

import "cmp"

// BinarySearch 二分查找
func BinarySearch[T cmp.Ordered](a []T, v T) int {
	if len(a) == 0 {
		return -1
	}

	l, r, m := 0, len(a)-1, 0

	for l <= r {
		m = (l + r) >> 1
		if a[m] == v {
			return m
		}
		if a[m] < v {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}
