package linear_structure

import "fmt"

// 顺序存储（数组）实现线性表

const capacity = 10

type list struct {
	value  [capacity]any
	length int
}

// 在给定位置前插入元素
func (l *list) insertBefore(i int, v any) {
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

// 查询给定位置的元素
func (l *list) valueAt(i int) any {
	if l.indexOutOfRange(i) {
		panic("索引越界")
	}
	return l.value[i]
}

// 查询给定元素第一次出现的位置，从未出现返回 -1 。
func (l *list) indexOf(v any) int {
	for i, d := range l.value {
		if d == v {
			return i
		}
	}
	return -1
}

// 删除指定位置的元素
func (l *list) deleteAt(i int) {
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
func (l *list) indexOutOfRange(i int) bool {
	return l.length == 0 || (i < 0 || i > l.length-1)
}

// 自描述
func (l *list) print() {
	fmt.Printf("length: %v\ndata: %v\n\n", l.length, l.value)
}

func ArrayLinerList() {
	// 初始化
	l := new(list)
	l.print()

	// 返回长度
	fmt.Printf("长度为 %v\n\n", l.length)

	// 插入
	l.insertBefore(0, 100)
	l.print()

	// 查询元素
	fmt.Printf("索引为 0 的元素为 %v\n\n", l.valueAt(0))

	// 查询索引
	fmt.Printf("元素 10 的索引为 %v\n\n", l.indexOf(10))
	fmt.Printf("元素 100 的索引为 %v\n\n", l.indexOf(100))

	// 删除
	l.deleteAt(0)
	l.print()
}
