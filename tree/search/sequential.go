package search

// SequentialSearch 顺序查找
func SequentialSearch(a []any, v any) int {
	for i := range a {
		if a[i] == v {
			return i
		}
	}
	return -1
}
