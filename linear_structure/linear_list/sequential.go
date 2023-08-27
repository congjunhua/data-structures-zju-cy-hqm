package linear_list

import "fmt"

// 顺序存储（数组）实现线性表

const capacity = 10

type List struct {
	Value  [capacity]any
	Length int
}

// InsertBefore 在给定位置前插入元素
func (l *List) InsertBefore(i int, v any) {
	switch l.Length {
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
		for p := l.Length - 1; p >= i; p-- {
			l.Value[p+1] = l.Value[p]
		}
	}
	l.Value[i] = v
	l.Length++
}

// ValueAt 查询给定位置的元素
func (l *List) ValueAt(i int) any {
	if l.indexOutOfRange(i) {
		panic("索引越界")
	}
	return l.Value[i]
}

// IndexOf 查询给定元素第一次出现的位置，从未出现返回 -1 。
func (l *List) IndexOf(v any) int {
	for i, d := range l.Value {
		if d == v {
			return i
		}
	}
	return -1
}

// DeleteAt 删除指定位置的元素
func (l *List) DeleteAt(i int) {
	if l.indexOutOfRange(i) {
		panic("索引越界")
	}
	for p := i; p <= l.Length-1; p++ {
		if p < l.Length-1 {
			l.Value[p] = l.Value[p+1]
			continue
		}
		l.Value[p] = nil
	}
	l.Length--
}

// 判断给定位置是否超出当前线性表的索引范围
func (l *List) indexOutOfRange(i int) bool {
	return l.Length == 0 || (i < 0 || i > l.Length-1)
}

// 自描述
func (l *List) print() {
	fmt.Printf("length: %v\ndata: %v\n\n", l.Length, l.Value)
}
