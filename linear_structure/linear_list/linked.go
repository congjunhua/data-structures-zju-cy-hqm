package linear_list

import (
	"errors"
	"fmt"
)

// LinkedList 链式存储（链表）实现线性表
type LinkedList struct {
	head *node
}

type node struct {
	value any
	next  *node
}

var (
	InvalidIndex = errors.New("位置超出了链表范围")
	NotExist     = errors.New("元素不存在")
)

// NewLinkedList 初始化
func NewLinkedList(values ...any) *LinkedList {
	ll := LinkedList{}
	if values != nil {
		ns := make([]*node, len(values))
		for i, d := range values {
			ns[i] = &node{
				value: d,
			}
		}
		for i := range ns {
			if i == len(ns)-1 {
				break
			}
			ns[i].next = ns[i+1]
		}
		ll.head = ns[0]
	}
	return &ll
}

// InsertAfter 在给定位置插入
func (ll *LinkedList) InsertAfter(i int, v any) error {
	if i < 0 || i > ll.Length()-1 {
		return InvalidIndex
	}
	t := ll.head
	for p := 0; p < i; p++ {
		t = t.next
	}
	t.next = &node{
		value: v,
		next:  t.next,
	}
	return nil
}

// Length 返回长度
func (ll *LinkedList) Length() int {
	l, c := 0, ll.head
	for c != nil {
		l++
		c = c.next
	}
	return l
}

// ValueAt 查询给定位置的元素
func (ll *LinkedList) ValueAt(i int) (any, error) {
	if i < 0 || i > ll.Length()-1 {
		return nil, InvalidIndex
	}
	c := ll.head
	for i > 0 {
		c = c.next
		i--
	}
	return c.value, nil
}

// IndexOf 查询给定元素第一次出现的位置，从未出现返回 -1 。
func (ll *LinkedList) IndexOf(v any) int {
	c, i := ll.head, 0
	for {
		if c.value == v {
			return i
		}
		if c.next == nil {
			break
		}
		c = c.next
		i++
	}
	return -1
}

// Delete 删除给定位置的节点
func (ll *LinkedList) Delete(i int) error {
	if i < 0 || i > ll.Length()-1 {
		return InvalidIndex
	}
	// 单独处理头节点
	if i == 0 {
		ll.head = ll.head.next
		return nil
	}
	// 定位非头节点的父节点
	p := ll.head
	for i-1 > 0 {
		p = p.next
		i--
	}
	p.next = p.next.next
	return nil
}

// Print 打印
func (ll *LinkedList) Print() {
	s := make([]node, 0)
	c := ll.head
	if c != nil {
		s = append(s, *c)
		for c.next != nil {
			c = c.next
			s = append(s, *c)
		}
	}
	fmt.Println(s)
}
