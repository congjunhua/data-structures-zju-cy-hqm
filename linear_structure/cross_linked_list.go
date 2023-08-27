package linear_structure

import (
	"fmt"
)

// 使用十字链表描述二维稀松数组。

// 链表
type crossLinkedList struct {
	rowHeads    []*rowHead
	columnHeads []*columnHead
	length      uint
}

// 行头
type rowHead struct {
	number uint
	right  *term
}

// 列头
type columnHead struct {
	number uint
	down   *term
}

// 节点
type term struct {
	row, column uint
	value       any
	right, down *term
}

// 初始化
func newCrossLinkedList(r, c uint) *crossLinkedList {
	rhs := make([]*rowHead, r)
	for i := uint(0); i < r; i++ {
		rhs[i] = &rowHead{number: i + 1}
	}
	chs := make([]*columnHead, c)
	for i := uint(0); i < c; i++ {
		chs[i] = &columnHead{number: i + 1}
	}
	return &crossLinkedList{
		rowHeads:    rhs,
		columnHeads: chs,
	}
}

// 插入节点
func (l *crossLinkedList) set(r, c uint, v any) error {
	if l.outOfRange(r, c) {
		return invalidIndex
	}

	t := &term{
		row:    r,
		column: c,
		value:  v,
	}

	// 行
	rh := l.rowHeads[r-1]
	if c == 1 {
		// 若为首列，则左侧链接行头，右侧（若有）链接子元素
		if rh.right != nil && rh.right.column > c {
			t.right = rh.right
		}
		rh.right = t
	} else {
		// 非首列，判断目标位置前是否存在元素
		pre := func() *term {
			if rh.right == nil {
				return nil
			}
			if rh.right.column >= c {
				return nil
			}
			p := rh.right
			for p.right != nil && p.right.column < c {
				p = p.right
			}
			return p
		}()
		if pre == nil {
			// 左侧不存在，则左侧链接行头
			rh.right = t
		} else {
			// 左侧存在，则左侧链接至该元素，该元素若存在列值大于目标列列值的子元素，则右侧链接该子元素
			if pre.right != nil && pre.right.column > c {
				t.right = pre.right
			}
			pre.right = t
		}
	}

	// 列
	ch := l.columnHeads[c-1]
	if r == 1 {
		// 若为首行，则上侧链接列头，右侧（若有）链接子元素
		if ch.down != nil && ch.down.row > r {
			t.down = ch.down
		}
		ch.down = t
	} else {
		// 非首行，判断目标位置前是否存在元素
		pre := func() *term {
			if ch.down == nil {
				return nil
			}
			if ch.down.row > r {
				return nil
			}
			p := ch.down
			for p.down != nil && p.down.row < r {
				p = p.down
			}
			return p
		}()
		if pre == nil {
			// 上侧不存在，则上侧链接至列头
			ch.down = t
		} else {
			// 上侧存在，则左侧链接至该元素，该元素若存在行值大于目标位置行值的子元素，则下侧链接该子元素
			if pre.down != nil && pre.down.row > r {
				t.down = pre.down
			}
			pre.down = t
		}
	}

	return nil
}

// 根据位置查询节点，若不存在，返回 notExist 错误。
func (l *crossLinkedList) get(r, c uint) (*term, error) {
	if l.outOfRange(r, c) {
		return nil, invalidIndex
	}
	var t *term
	rh := l.rowHeads[r-1]
	if rh.right == nil {
		return nil, notExist
	}
	t = rh.right
	for t.right != nil && t.column < c {
		t = t.right
	}
	if t.column != c {
		return nil, notExist
	}
	return t, nil
}

// 删除
func (l *crossLinkedList) delete(r, c uint) error {
	if l.outOfRange(r, c) {
		return invalidIndex
	}

	t, e := l.get(r, c)
	if e != nil {
		return e
	}

	// 行
	rh := l.rowHeads[r-1]
	left := rh.right
	for left.right != nil && left.right.column < c {
		left = left.right
	}
	if left.column >= c {
		left = nil
	}
	if left == nil {
		// 行链不存在父节点，则调整行头的右侧链接
		rh.right = t.right
	} else {
		// 行链存在父节点，则调整父节点的右侧链接
		left.right = t.right
	}

	// 列
	ch := l.columnHeads[c-1]
	up := ch.down
	for up.down != nil && up.down.row < r {
		up = up.down
	}
	if up.row >= r {
		up = nil
	}
	if up == nil {
		// 列链不存在父节点，则调整列头的下侧链接
		ch.down = t.down
	} else {
		// 列链存在父节点，则调整父节点的下侧链接
		up.down = t.down
	}

	return nil
}

// 打印
func (l *crossLinkedList) print() {
	fmt.Println(l)
	m := make([][]any, len(l.rowHeads))
	for i := 0; i < len(m); i++ {
		m[i] = make([]any, len(l.columnHeads))
	}
	for i, rh := range l.rowHeads {
		t := rh.right
		for t != nil {
			m[i][t.column-1] = t
			t = t.right
		}
	}
	for i := 0; i < len(m); i++ {
		for _, t := range m[i] {
			s := fmt.Sprintf("%v", t)
			for len(s) < 36 {
				s += " "
			}
			fmt.Print(s)
		}
		fmt.Println()
	}
}

// 判断位置是否正确
func (l *crossLinkedList) outOfRange(r, c uint) bool {
	return r < 1 || r > uint(len(l.rowHeads)) || c < 1 || c > uint(len(l.columnHeads))
}
