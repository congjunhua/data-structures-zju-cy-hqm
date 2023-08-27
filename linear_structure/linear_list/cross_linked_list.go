package linear_list

import (
	"fmt"
)

// 使用十字链表描述二维稀松数组。

// CrossLinkedList 链表
type CrossLinkedList struct {
	RowHeads    []*RowHead
	ColumnHeads []*ColumnHead
	Length      uint
}

// RowHead 行头
type RowHead struct {
	Number uint
	Right  *Term
}

// ColumnHead 列头
type ColumnHead struct {
	Number uint
	Down   *Term
}

// Term 节点
type Term struct {
	Row, Column uint
	Value       any
	Right, Down *Term
}

// NewCrossLinkedList 初始化
func NewCrossLinkedList(r, c uint) *CrossLinkedList {
	rhs := make([]*RowHead, r)
	for i := uint(0); i < r; i++ {
		rhs[i] = &RowHead{Number: i + 1}
	}
	chs := make([]*ColumnHead, c)
	for i := uint(0); i < c; i++ {
		chs[i] = &ColumnHead{Number: i + 1}
	}
	return &CrossLinkedList{
		RowHeads:    rhs,
		ColumnHeads: chs,
	}
}

// Set 插入
func (l *CrossLinkedList) Set(r, c uint, v any) error {
	if l.outOfRange(r, c) {
		return InvalidIndex
	}

	t := &Term{
		Row:    r,
		Column: c,
		Value:  v,
	}

	// 行
	rh := l.RowHeads[r-1]
	if c == 1 {
		// 若为首列，则左侧链接行头，右侧（若有）链接子元素
		if rh.Right != nil && rh.Right.Column > c {
			t.Right = rh.Right
		}
		rh.Right = t
	} else {
		// 非首列，判断目标位置前是否存在元素
		pre := func() *Term {
			if rh.Right == nil {
				return nil
			}
			if rh.Right.Column >= c {
				return nil
			}
			p := rh.Right
			for p.Right != nil && p.Right.Column < c {
				p = p.Right
			}
			return p
		}()
		if pre == nil {
			// 左侧不存在，则左侧链接行头
			rh.Right = t
		} else {
			// 左侧存在，则左侧链接至该元素，该元素若存在列值大于目标列列值的子元素，则右侧链接该子元素
			if pre.Right != nil && pre.Right.Column > c {
				t.Right = pre.Right
			}
			pre.Right = t
		}
	}

	// 列
	ch := l.ColumnHeads[c-1]
	if r == 1 {
		// 若为首行，则上侧链接列头，右侧（若有）链接子元素
		if ch.Down != nil && ch.Down.Row > r {
			t.Down = ch.Down
		}
		ch.Down = t
	} else {
		// 非首行，判断目标位置前是否存在元素
		pre := func() *Term {
			if ch.Down == nil {
				return nil
			}
			if ch.Down.Row > r {
				return nil
			}
			p := ch.Down
			for p.Down != nil && p.Down.Row < r {
				p = p.Down
			}
			return p
		}()
		if pre == nil {
			// 上侧不存在，则上侧链接至列头
			ch.Down = t
		} else {
			// 上侧存在，则左侧链接至该元素，该元素若存在行值大于目标位置行值的子元素，则下侧链接该子元素
			if pre.Down != nil && pre.Down.Row > r {
				t.Down = pre.Down
			}
			pre.Down = t
		}
	}

	return nil
}

// Get 根据位置查询，若不存在，返回 NotExist 错误。
func (l *CrossLinkedList) Get(r, c uint) (*Term, error) {
	if l.outOfRange(r, c) {
		return nil, InvalidIndex
	}
	var t *Term
	rh := l.RowHeads[r-1]
	if rh.Right == nil {
		return nil, NotExist
	}
	t = rh.Right
	for t.Right != nil && t.Column < c {
		t = t.Right
	}
	if t.Column != c {
		return nil, NotExist
	}
	return t, nil
}

// Delete 删除
func (l *CrossLinkedList) Delete(r, c uint) error {
	if l.outOfRange(r, c) {
		return InvalidIndex
	}

	t, e := l.Get(r, c)
	if e != nil {
		return e
	}

	// 行
	rh := l.RowHeads[r-1]
	left := rh.Right
	for left.Right != nil && left.Right.Column < c {
		left = left.Right
	}
	if left.Column >= c {
		left = nil
	}
	if left == nil {
		// 行链不存在父节点，则调整行头的右侧链接
		rh.Right = t.Right
	} else {
		// 行链存在父节点，则调整父节点的右侧链接
		left.Right = t.Right
	}

	// 列
	ch := l.ColumnHeads[c-1]
	up := ch.Down
	for up.Down != nil && up.Down.Row < r {
		up = up.Down
	}
	if up.Row >= r {
		up = nil
	}
	if up == nil {
		// 列链不存在父节点，则调整列头的下侧链接
		ch.Down = t.Down
	} else {
		// 列链存在父节点，则调整父节点的下侧链接
		up.Down = t.Down
	}

	return nil
}

// Print 打印
func (l *CrossLinkedList) Print() {
	fmt.Println(l)
	m := make([][]any, len(l.RowHeads))
	for i := 0; i < len(m); i++ {
		m[i] = make([]any, len(l.ColumnHeads))
	}
	for i, rh := range l.RowHeads {
		t := rh.Right
		for t != nil {
			m[i][t.Column-1] = t
			t = t.Right
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
func (l *CrossLinkedList) outOfRange(r, c uint) bool {
	return r < 1 || r > uint(len(l.RowHeads)) || c < 1 || c > uint(len(l.ColumnHeads))
}
