package linear_structure

import (
	"errors"
	"fmt"
	"log"
)

// 链式存储（链表）实现线性表

type node struct {
	value any
	next  *node
}

type linkedList struct {
	head *node
}

var (
	invalidIndex = errors.New("位置超出了链表范围")
	notExist     = errors.New("元素不存在")
)

// 初始化
func newLinkedList(values ...any) *linkedList {
	ll := linkedList{}
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

// 在给定位置插入
func (ll *linkedList) insertAfter(i int, v any) error {
	if i < 0 || i > ll.length()-1 {
		return invalidIndex
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

// 长度
func (ll *linkedList) length() int {
	l, c := 0, ll.head
	for c != nil {
		l++
		c = c.next
	}
	return l
}

// 查询给定位置的元素
func (ll *linkedList) valueAt(i int) (any, error) {
	if i < 0 || i > ll.length()-1 {
		return nil, invalidIndex
	}
	c := ll.head
	for i > 0 {
		c = c.next
		i--
	}
	return c.value, nil
}

// 查询给定元素第一次出现的位置，从未出现返回 -1 。
func (ll *linkedList) indexOf(v any) int {
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

// 删除给定位置的节点
func (ll *linkedList) delete(i int) error {
	if i < 0 || i > ll.length()-1 {
		return invalidIndex
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

// 自描述
func (ll *linkedList) print() {
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

func LinkedListLinerList() {
	// 初始化
	ll := newLinkedList(1)
	ll.print()

	// 插入
	if e := ll.insertAfter(0, 2); e != nil {
		log.Println(e)
	}
	ll.print()
	if e := ll.insertAfter(1, 3); e != nil {
		log.Println(e)
	}
	ll.print()

	// 长度
	fmt.Println("长度 = ", ll.length())

	// 查询给定位置的元素
	for i := 0; i < ll.length(); i++ {
		fmt.Println(ll.valueAt(i))
	}

	// 查询给定元素第一次出现的位置
	for _, v := range []any{1, 2, 3} {
		fmt.Println(ll.indexOf(v))
	}

	// 删除给定位置的节点
	fmt.Println(ll.delete(2))
	ll.print()
	fmt.Println(ll.delete(1))
	ll.print()
	fmt.Println(ll.delete(0))
	ll.print()
}
