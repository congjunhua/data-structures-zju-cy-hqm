package linear_list

import (
	"errors"
	"fmt"
)

// 链式存储（链表）实现线性表

type Node struct {
	Value any
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

var (
	InvalidIndex = errors.New("位置超出了链表范围")
	NotExist     = errors.New("元素不存在")
)

// NewLinkedList 初始化
func NewLinkedList(values ...any) *LinkedList {
	ll := LinkedList{}
	if values != nil {
		ns := make([]*Node, len(values))
		for i, d := range values {
			ns[i] = &Node{
				Value: d,
			}
		}
		for i := range ns {
			if i == len(ns)-1 {
				break
			}
			ns[i].Next = ns[i+1]
		}
		ll.Head = ns[0]
	}
	return &ll
}

// InsertAfter 在给定位置插入
func (ll *LinkedList) InsertAfter(i int, v any) error {
	if i < 0 || i > ll.Length()-1 {
		return InvalidIndex
	}
	t := ll.Head
	for p := 0; p < i; p++ {
		t = t.Next
	}
	t.Next = &Node{
		Value: v,
		Next:  t.Next,
	}
	return nil
}

// Length 返回长度
func (ll *LinkedList) Length() int {
	l, c := 0, ll.Head
	for c != nil {
		l++
		c = c.Next
	}
	return l
}

// ValueAt 查询给定位置的元素
func (ll *LinkedList) ValueAt(i int) (any, error) {
	if i < 0 || i > ll.Length()-1 {
		return nil, InvalidIndex
	}
	c := ll.Head
	for i > 0 {
		c = c.Next
		i--
	}
	return c.Value, nil
}

// IndexOf 查询给定元素第一次出现的位置，从未出现返回 -1 。
func (ll *LinkedList) IndexOf(v any) int {
	c, i := ll.Head, 0
	for {
		if c.Value == v {
			return i
		}
		if c.Next == nil {
			break
		}
		c = c.Next
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
		ll.Head = ll.Head.Next
		return nil
	}
	// 定位非头节点的父节点
	p := ll.Head
	for i-1 > 0 {
		p = p.Next
		i--
	}
	p.Next = p.Next.Next
	return nil
}

// Print 打印
func (ll *LinkedList) Print() {
	s := make([]Node, 0)
	c := ll.Head
	if c != nil {
		s = append(s, *c)
		for c.Next != nil {
			c = c.Next
			s = append(s, *c)
		}
	}
	fmt.Println(s)
}
