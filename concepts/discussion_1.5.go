package concepts

/*
	讨论1.5 分析“二分法”
	https://www.icourse163.org/learn/ZJU-93001?tid=1470659487#/learn/content?type=detail&id=1254705091&cid=1286093701
*/

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type float interface {
	~float32 | ~float64
}

type number interface {
	signed | unsigned | float
}

// BinarySearch 返回 target 在 arr 中的 index ，若不存在，返回 -1 。
func BinarySearch[T number](arr []T, target T) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
