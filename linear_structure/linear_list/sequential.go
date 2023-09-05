package linear_list

import "fmt"

const capacity = 10

// SequentialList 顺序存储（数组）实现线性表
type SequentialList struct {
	value  [capacity]any
	length int
}

// InsertBefore 在给定位置前插入元素
func (l *SequentialList) InsertBefore(i int, v any) {
	switch l.length {
	case capacity:
		panic("表已满")
	case 0:
		if i != 0 {
			panic("空表只能从从表头插入")
		}
	default:
		if l.indexOutOfRange(i) {
			panic("索引越界")
		}
		for p := l.length - 1; p >= i; p-- {
			l.value[p+1] = l.value[p]
		}
	}
	l.value[i] = v
	l.length++
}

// ValueAt 查询给定位置的元素
func (l *SequentialList) ValueAt(i int) any {
	if l.indexOutOfRange(i) {
		panic("索引越界")
	}
	return l.value[i]
}

// IndexOf 查询给定元素第一次出现的位置，从未出现返回 -1 。
func (l *SequentialList) IndexOf(v any) int {
	for i, d := range l.value {
		if d == v {
			return i
		}
	}
	return -1
}

// DeleteAt 删除指定位置的元素
func (l *SequentialList) DeleteAt(i int) {
	if l.indexOutOfRange(i) {
		panic("索引越界")
	}
	for p := i; p <= l.length-1; p++ {
		if p < l.length-1 {
			l.value[p] = l.value[p+1]
			continue
		}
		l.value[p] = nil
	}
	l.length--
}

// 判断给定位置是否超出当前线性表的索引范围
func (l *SequentialList) indexOutOfRange(i int) bool {
	return l.length == 0 || (i < 0 || i > l.length-1)
}

// 自描述
func (l *SequentialList) print() {
	fmt.Printf("length: %v\ndata: %v\n\n", l.length, l.value)
}
