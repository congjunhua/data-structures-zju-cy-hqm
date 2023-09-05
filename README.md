# 数据结构（陈越、何钦铭）

# 一、基本概念

## 1.1 数据结构

计算机存储、组织数据的方式。

## 1.2 算法

### 1.2.1 定义

接受一些可能的输入，在有限的计算步骤后，输出特定结果的指令集。

其中，指令集中的每条指令须满足以下要求：

- 目标明确、无歧义

- 不能超过计算机的计算能力上限

- 不依赖具体的语言

- 不依赖具体的实现方式

### 1.2.2 评价

空间复杂度 $S(n)$ ：算法执行时占用的存储空间。

时间复杂度 $T(n)$ ：算法的执行耗时。

最坏情况复杂度 $T_{worst}(n)$

最佳情况复杂度 $T_{best}(n)$

平均复杂度 $T_{avg}(n)$

### 1.2.3 渐进表示法

$T(n)=O(f(n))$

存在常数 $C>0,n_0>0$ 使得 $n≥n_0$ 时有 $T(n)≤C·f(n)$ ，即对于充分大的 $n$ 而言， $f(n)$ 表示 $T(n)$ 的某种上限。

$T(n)=Ω(g(n))$

存在常数 $C>0,n_0>0$ 使得 $n≥n_0$ 时有 $T(n)≥C·g(n)$ ，即对于充分大的 $n$ 而言， $g(n)$ 表示 $T(n)$ 的某种下限。

$T(n)=θ(h(n))$

同时存在 $T(n)=O(h(n))$ 和 $T(n)=Ω(h(n))$ ，即 $h(n)$ 既是 $T(n)$ 的上限也是 $T(n)$ 的下限。

**注意**

- 当我们在关注复杂度上下限的时候，往往是关注 $O(n)$ 的最小值和 $Ω(n)$ 的最大值，因为过大的上限和过小的下限对于分析算法的复杂度没多少帮助。

- 对于 $log_xn$ ，实践中可简化为 $log\ n$，因为底数对复杂度的影响基本可以忽略。

- $log\ n$ < $n$ < $n\ log\ n$ < $n^2$ < $2^n$ ，复杂度大于 $n\ log\ n$ 的算法不是好算法，应本能地尝试降低复杂度。

**如何分析复杂度？**

若两个算法分别具有复杂度 $T_1(n)=O(f_1(n))$ 和复杂度 $T_2(n)=O(f_2(n))$，则：

- $T_1(n)+T_2(n)=max(O(f_1(n)), O(f_2(n)))$

- $T_1(n)·T_2(n)=O(f_1(n)·f_2(n))$

基于上述规则，可知：

- 对于 $n$ 的 $k$ 阶多项式， $T(n)=θ(n^k)$ 。

- 对于循环，复杂度等于次数乘以循环体的复杂度。

- 对于条件，复杂度取分支中复杂度的最大值。

## 1.3 最大子列和

**暴力计算**

```go
func calc(nums []int) int {
	l, s, m := len(nums), 0, 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l+1; j++ {
			s = 0
			for _, v := range nums[i:j] {
				s += v
			}
			m = max(m, s)
		}
	}
	return m
}
```

**消除第三层循环后的优化版本**

```go
func better(nums []int) int {
	l, s, m := len(nums), 0, 0
	for i := 0; i < l; i++ {
		s = 0
		for j := i; j < l; j++ {
			s += nums[j]
			m = max(m, s)
		}
	}
	return m
}
```

**分治法**

```go
func divideAndConquer(a []int) int {
	return maxSum(a, 0, len(a)-1)
}

func maxSum(a []int, l, r int) int {
	// 若子数组为单个元素，则直接返回该元素
	if l == r {
		return a[l]
	}

	// 中间元素索引
	m := (l + r) >> 1

	// 左区
	left := maxSum(a, l, m)

	// 右区
	right := maxSum(a, m+1, r)

	// 跨区
	sum, acrossLeft, acrossRight := 0, 0, 0
	for i := m; i >= l; i-- {
		sum += a[i]
		acrossLeft = max(acrossLeft, sum)
	}
	sum = 0
	for i := m + 1; i <= r; i++ {
		sum += a[i]
		acrossRight = max(acrossRight, sum)
	}
	across := acrossLeft + acrossRight

	return max(left, right, across)
}
```

**在线处理**

```go
func onlineProcess(a []int) int {
	m := a[0]
	for i := 1; i < len(a); i++ {
		if a[i]+a[i-1] > a[i] {
			a[i] += a[i-1]
		}
		if a[i] > m {
			m = a[i]
		}
		m = max(m, a[i])
	}
	return m
}
```

详见`/concepts/max_subarray.go`。



# 二、线性结构

## 2.1 线性表

### 2.1.1 多项式

如何使用程序语言描述 $f(x) = a_0 + a_1 x + a_2 x^2 + \ldots + a_{n-1} x^{n-1} + a_n x^n$ ？

核心字段：

- 项数 $n$

- 系数 $a_i$

- 指数 $i$

**方法1：顺序存储结构**

即使用一个数组表示多项式，多项式的项与数组的项一一对应，项树 $n$ 对应数组长度，系数 $a_i$ 对应数组第 `i` 个元素的值，指数 $i$ 对应数组的索引。

例如 $f(x)=4x^5-3x^2+1$ 可以使用 `[1, 0, -3, 0, 0, 4]` 表示。

这种表示法很直观，但显而易见的问题是对存储空间的浪费，因为系数为零的项完全不需要表示。

**方法2：仅包含非零项的顺序存储结构**

方法 1 中的函数可以使用 `[(4, 5), (-3, 2), (1, 0)]` 表示。

**方法3：链表**

```go
type node struct {
    coefficient int
    exponent    int
    next        *node
}
```

### 2.1.2 线性表

**线性表**（Linear List）是由 $n$（$n≥0$）个元素组成的有限序列。

其中：

- $n$ 为线性表的**长度**。

- $n$ 为零的线性表称为**空表**。

- 起始位置称为**表头**，结束位置称为**表尾**。

- 非空的线性表可以表示为 $a_0,\ a_1,\ ...\ a_{n-2},\ a_{n-1}$ 。

**数据集**： $a_0,\ a_1,\ ...\ a_{n-2},\ a_{n-1}$ 。

**方法集**：

- 初始化

- 在给定位置插入元素

- 返回长度

- 查询给定位置的元素

- 查询给定元素第一次出现的位置

- 删除给定位置的元素

### 2.1.3 顺序表

使用**数组**实现线性表。

```go
const capacity = 100

type list struct {
    value   [capacity]any
    length int
}
```

其中：

- `capacity` 为数组的长度（线性表的最大长度）。

- `length` 为线性表的长度。

对于表中任意一项 $a_i$ ，满足 $0≤i≤length-1≤capacity-1$ 。

```go
const capacity = 10

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
```

详见 `/linear_structure/linear_list/sequential.go` 。

### 2.1.4 链表

使用**链表**实现线性表。

数组实现的线性表，元素之间不仅逻辑上相邻，物理（存储介质）上也是相邻的，后者导致在插入和删除时，须手动更新受影响元素的物理位置。

而链表仅仅在逻辑上相邻，进而在插入和删除时不再需要移动受影响的元素，仅需修改链接即可。

```go
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
```

详见 `/linear_structure/linear_list/chained.go` 。

### 2.1.5 广义表

如何表示 $P(x,\ y)=9x^{12}y^{2}+4x^{12}+15x^8y^3-x^8y+3x^2$ 这个二元多项式？

思路：简化为 $x$ 的一元多项式：

1. $p(x,\ y)=(9y^2+4)x^{12}+(15y^3-y)x^8+3x^2$

2. $p(x)=ax^{12}+bx^8+cx^2$

其中， $a,\ b,\ c$ 不再是常数，而是 $y$ 的一元多项式。

```go
type y struct {
    constant int
    power    uint
    next     *y
}

type x[t int | *y] struct {
    constant t
    power    uint
    next     *x[t]
}
```

这种对线性表的扩展，就是**广义表**：

$LS=(a_1,\ a_2,\ ...\ a_{n-1},\ a_n)$

其中：

- $n$ 为长度。

- $a_i$ 可以是常数，也可以是广义表。

- 对于非空的广义表， $a_1$ 表示表头，剩余的所有元素组成的表为表尾。

除非所有元素都是原子，否则广义表表现为非线形。

### 2.1.6 多重链表

若节点指向多个不同的链表，则该节点所在的链表称为**多重链表**。

**问题**：如何描述一个矩阵？

首先想到二维数组，但二维数组缺点明显：大小需要事先确定、稀松矩阵会造成存储空间的浪费。

这种情况下，可以使用多重链表中的**十字链表**来实现。

```go
type CrossLinkedList struct {
	RowHeads    []*rowHead
	ColumnHeads []*columnHead
	Length      uint
}

type rowHead struct {
	number uint
	right  *Term
}

type columnHead struct {
	number uint
	down   *Term
}

type Term struct {
	Row, Column uint
	Value       any
	Right, Down *Term
}

// NewCrossLinkedList 初始化
func NewCrossLinkedList(r, c uint) *CrossLinkedList {
	rhs := make([]*rowHead, r)
	for i := uint(0); i < r; i++ {
		rhs[i] = &rowHead{number: i + 1}
	}
	chs := make([]*columnHead, c)
	for i := uint(0); i < c; i++ {
		chs[i] = &columnHead{number: i + 1}
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
		if rh.right != nil && rh.right.Column > c {
			t.Right = rh.right
		}
		rh.right = t
	} else {
		// 非首列，判断目标位置前是否存在元素
		pre := func() *Term {
			if rh.right == nil {
				return nil
			}
			if rh.right.Column >= c {
				return nil
			}
			p := rh.right
			for p.Right != nil && p.Right.Column < c {
				p = p.Right
			}
			return p
		}()
		if pre == nil {
			// 左侧不存在，则左侧链接行头
			rh.right = t
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
		if ch.down != nil && ch.down.Row > r {
			t.Down = ch.down
		}
		ch.down = t
	} else {
		// 非首行，判断目标位置前是否存在元素
		pre := func() *Term {
			if ch.down == nil {
				return nil
			}
			if ch.down.Row > r {
				return nil
			}
			p := ch.down
			for p.Down != nil && p.Down.Row < r {
				p = p.Down
			}
			return p
		}()
		if pre == nil {
			// 上侧不存在，则上侧链接至列头
			ch.down = t
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
	if rh.right == nil {
		return nil, NotExist
	}
	t = rh.right
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
	left := rh.right
	for left.Right != nil && left.Right.Column < c {
		left = left.Right
	}
	if left.Column >= c {
		left = nil
	}
	if left == nil {
		// 行链不存在父节点，则调整行头的右侧链接
		rh.right = t.Right
	} else {
		// 行链存在父节点，则调整父节点的右侧链接
		left.Right = t.Right
	}

	// 列
	ch := l.ColumnHeads[c-1]
	up := ch.down
	for up.Down != nil && up.Down.Row < r {
		up = up.Down
	}
	if up.Row >= r {
		up = nil
	}
	if up == nil {
		// 列链不存在父节点，则调整列头的下侧链接
		ch.down = t.Down
	} else {
		// 列链存在父节点，则调整父节点的下侧链接
		up.Down = t.Down
	}

	return nil
}
```

详见 `/linear_structure/linear_list/cross_linked_list.go`。

## 2.2 堆栈

### 2.2.1 概念

问题：计算机如何对表达式 $5+6/2-3*4$ 进行求值？

因为运算符优先级的存在，导致在看到运算符后面的数字时，并不难判断是否要立即执行运算，因为后面的运算符的优先级可能会高于当前的运算符，只有当遍历所有运算符之后，才能判断运算的执行顺序。

要实现看到运算符就立即执行计算，可以将运算符放到两个数字之后，而不是放在两个数字之间，比如

$a+b*c-d/e$

可以写成如下形式：

$abc*+de/-$

这种写法称为**后缀表达式**，而普通的写法称为中缀表达式。

如果将运算符放在两数之前，则上述表达式可以写成 $-+a*bc/de$ ，此为前缀表达式。

所以，对于上述问题：

1. 计算机首先将中缀表达式转化为后缀表达式 $562/+34*-$

2. 然后从左往右扫描该后缀表达式

3. 遇到数字暂存

4. 遇到运算符对暂存的数字的后两个数字执行运算

5. 直到表达式遍历完成

可以发现，需要一种满足「**后入先出**」的数据结构来存储待运算的数字，这种数据结构，即为**堆栈**（stack）。

其中：

- 只能在栈的一端插入或者删除，这个端点称之为**栈顶**

- 插入数据称为**入栈**（Push）

- 删除数据称为**出栈**（Pop）

- 后入先出简称 **LIFO**（Last In First Out）

**数据集**：零个或多个元素的有限线性表。

**方法集**：

- 初始化一个空栈

- 判断是否为空

- 判断是否已满

- 查询栈顶元素

- 元素入栈

- 元素出栈

### 2.2.2 顺序存储

```go
// 使用数组实现堆栈

type ArrayStack struct {
	Values [capacity]any
	Top    int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		Top: -1,
	}
}

func (s *ArrayStack) Empty() bool {
	return s.Top == -1
}

func (s *ArrayStack) Full() bool {
	return s.Top == len(s.Values)-1
}

func (s *ArrayStack) Peek() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	return s.Values[s.Top], nil
}

func (s *ArrayStack) Push(v any) error {
	if s.Full() {
		return FullError
	}
	s.Top++
	s.Values[s.Top] = v
	return nil
}

func (s *ArrayStack) Pop() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	v := s.Values[s.Top]
	s.Top--
	return v, nil
}
```

详见 `/linear_structure/stack/array.go` 。

```go
// 使用切片实现堆栈

type SliceStack struct {
	Values []any
}

func NewSliceStack() *SliceStack {
	return &SliceStack{}
}

func (s *SliceStack) Empty() bool {
	return len(s.Values) == 0
}

// 因为切片是动态数组，不关心栈满。
// func (s *SliceStack) Full() bool {}

func (s *SliceStack) Peek() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	return s.Values[len(s.Values)-1], nil
}

func (s *SliceStack) Push(v any) {
	s.Values = append(s.Values, v)
}

func (s *SliceStack) Pop() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	i := len(s.Values) - 1
	v := s.Values[i]
	s.Values = s.Values[:i]
	return v, nil
}
```

详见 `/linear_structure/stack/slice.go` 。

### 2.2.3 链式存储

使用**链表**实现堆栈。

```go
type LinkedListStack struct {
	top *node
}

type node struct {
	value any
	next  *node
}

func (s *LinkedListStack) Empty() bool {
	return s.top == nil
}

func (s *LinkedListStack) Peek() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	return s.top.value, nil
}

func (s *LinkedListStack) Push(v any) {
	s.top = &node{value: v, next: s.top}
}

func (s *LinkedListStack) Pop() (any, error) {
	if s.Empty() {
		return nil, EmptyError
	}
	v := s.top.value
	s.top = s.top.next
	return v, nil
}
```

详见 `/linear_structure/stack/linked_list.go` 。

## 2.3 队列

### 2.3.1 概念

队列是一特殊的线性表，具有以下特征：

- 只能在**队尾**（rear）插入

- 只能在**队头**（front）删除

- 插入数据称为**入队**

- 删除数据称为**出队**

- 先进先出 **FIFO**（First In First Out）

**数据集**：零个或多个元素的有限线性表。

**方法集**：

- 初始化一个空对列
- 判断是否为空
- 判断是否已满
- 查询队头元素
- 元素入队
- 元素出队

### 2.3.2 顺序存储

使用**数组**实现循环对列。

```go
const capacity = 4

var (
	EmptyQueue = errors.New("对列为空")
	FullQueue  = errors.New("对列已满")
)

// ArrayQueue 使用数组实现循环队列
type ArrayQueue struct {
	Values      [capacity]any
	Empty       bool
	Front, Rear uint
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		Empty: true,
	}
}

func (q *ArrayQueue) Full() bool {
	return next(q.Rear) == q.Front
}

func (q *ArrayQueue) Peek() (any, error) {
	if q.Empty {
		return nil, EmptyQueue
	}
	return q.Values[q.Front], nil
}

func (q *ArrayQueue) Put(v any) error {
	if q.Full() {
		return FullQueue
	}
	if !q.Empty {
		q.Rear++
	}
	q.Values[q.Rear] = v
	if q.Empty {
		q.Empty = !q.Empty
	}
	return nil
}

func (q *ArrayQueue) Poll() (any, error) {
	if q.Empty {
		return nil, EmptyQueue
	}
	v := q.Values[q.Front]
	q.Values[q.Front] = nil
	if q.Front != q.Rear {
		q.Front = next(q.Front)
	} else {
		q.Empty = !q.Empty
	}
	return v, nil
}

func next(p uint) uint {
	return (p + 1) % capacity
}
```

详见 `/linear_structure/queue/array.go` 。

### 2.3.3 链式存储

使用**链表**实现对列。

```go
type LinkedListQueue struct {
	front, rear *node
}

type node struct {
	value any
	next  *node
}

func (l *LinkedListQueue) Empty() bool {
	return l.front == nil
}

func (l *LinkedListQueue) Peek() (any, error) {
	if l.Empty() {
		return nil, EmptyQueue
	}
	return l.front.value, nil
}

func (l *LinkedListQueue) Put(v any) {
	n := &node{
		value: v,
	}

	if l.Empty() {
		l.front = n
		l.rear = n
		return
	}

	l.rear.next = n
	l.rear = n
}

func (l *LinkedListQueue) Poll() (any, error) {
	if l.Empty() {
		return nil, EmptyQueue
	}

	v := l.front.value

	if l.front == l.rear {
		l.front = nil
		l.rear = nil
		return v, nil
	}

	nt := l.front.next
	l.front.next = nil
	l.front = nt

	return v, nil
}
```

详见 `/linear_structure/queue/linked_list.go` 。

### 2.3.4 计算多项式

$p1=3x^5+4x^4-x^3+2x-1$

$p2=2x^4+x^3-7x^2+x$

求相加后的多项式 $p3$ 和相乘后的多项式 $p4$ 。

```go
type Polynomial indeterminate

type indeterminate struct {
	coefficient int
	power       uint
	next        *indeterminate
}

func NewPolynomial(vs []indeterminate) *Polynomial {
	for i := 0; i < len(vs)-1; i++ {
		vs[i].next = &vs[i+1]
	}
	v := Polynomial{}
	if len(vs) > 0 {
		v = Polynomial(vs[0])
	}
	return &v
}

func (p *Polynomial) Add(another *Polynomial) *Polynomial {
	m := map[uint]int{}

	ind := (*indeterminate)(p)
	for ind != nil {
		m[p.power] = p.coefficient
		ind = ind.next
	}

	ind = (*indeterminate)(another)
	for ind != nil {
		if _, ok := m[ind.power]; ok {
			m[ind.power] += ind.coefficient
		}
		ind = ind.next
	}

	vs := make([]indeterminate, 0)
	for power, coefficient := range m {
		vs = append(
			vs, indeterminate{
				coefficient: coefficient,
				power:       power,
			},
		)
	}

	sort.Slice(
		vs, func(i, j int) bool {
			return vs[i].power > vs[j].power
		},
	)

	return NewPolynomial(vs)
}

func (p *Polynomial) MultipliedBy(another *Polynomial) *Polynomial {
	if p == nil || another == nil {
		return nil
	}
	m := map[uint]int{}
	pi := (*indeterminate)(p)
	for pi != nil {
		ai := (*indeterminate)(another)
		for ai != nil {
			coefficient := pi.coefficient * ai.coefficient
			power := pi.power + ai.power
			if _, ok := m[power]; ok {
				m[power] += coefficient
			} else {
				m[power] = coefficient
			}
			ai = ai.next
		}
		pi = pi.next
	}
	vs := make([]indeterminate, 0)
	for power, coefficient := range m {
		vs = append(
			vs, indeterminate{
				coefficient: coefficient,
				power:       power,
			},
		)
	}
	sort.Slice(
		vs, func(i, j int) bool {
			return vs[i].power > vs[j].power
		},
	)
	return NewPolynomial(vs)
}
```

详见 `/linear_structure/queue/polynomial.go` 。



# 三、树

## 3.1 树与树的表示

### 3.1.1 顺序查找

即遍历数组。

```go
func SequentialSearch(a []any, v any) int {
	for i := range a {
		if a[i] == v {
			return i
		}
	}
	return -1
}
```

时间复杂度 $O(n)$ ，详见 `/tree/search/sequential.go` 。

### 3.1.2 二分查找

首先升序对数组进行排序，然后进入循环，每次循环，只对中间位置的元素与目标元素进行相等性和排序性比较：

1. 若中间位置的元素与目标元素相等，直接返回中间元素的下标；

2. 若中间位置的元素小于目标元素，说明目标元素必不可能出现在中间位置元素的左侧，待搜索区域缩小一半；

3. 若中间位置的元素大于目标元素，说明目标元素必不可能出现在中间位置元素的右侧，待搜索区域减小一半；

直至待搜索区域为零。

```go
func BinarySearch[T cmp.Ordered](a []T, v T) int {
	if len(a) == 0 {
		return -1
	}

	l, r, m := 0, len(a)-1, 0

	for l <= r {
		m = (l + r) >> 1
		if a[m] == v {
			return m
		}
		if a[m] < v {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}
```

时间复杂度 $O(log_2\ n)$ ，详见 `/tree/search/binary.go` 。

### 3.1.3 概念

由有限节点构成的具有层次关系的集合。

其中：

- 没有父节点的节点称为根（Root）节点

- 每个节点都有 $n(n≥0)$ 个有限子节点

- 每个非根节点有且仅有一个父节点

- 每个非根节点可分为多个不相交的子树

- 没有环路（cycle）

- 一棵 $n$ 个节点的树有 $n-1$ 个边（链接）

术语：

- **节点的度**（Degree）：节点的子树棵数。

- **树的度**（Degree of tree）：所有节点中节点度的最大值。

- **叶节点**（Leaf）：度为零的节点。

- **层**（Level）：根节点为 $1$ 层，叶节点为最后一层，亦称为**深度**。

- **距离**（Distance）：两个节点间最短路径的边数。

## 3.2 二叉树

### 3.2.1 概念

由度为 $n(0≤n≤2)$ 的节点组成的树，称为**二叉树**。

其中：

- 第 $i(i≥1)$ 层的最大节点数为 $2^{i-1}$
- 深度为 $d(d≥1)$ 的二叉树最大节点数为 $2^d-1$
- 叶节点总数永远比度为 $2$ 的节点的总数大 $1$

几个特殊的二叉树：

- **斜二叉树**：所有节点有且仅有左子树或所有节点有且仅有右子树。

- **完全二叉树**：深度为 $d$ ， $d$ 层从左往右除最后一个节点之外所有节点的父节点的度为 $2$ ， $d$ 层外的每一层的所有节点的父节点的度为 $2$ 。

- **满二叉树**：深度为 $d$ ， $d$ 层外每一层所有节点的度为 $2$ 。

### 3.2.2 数据结构

**数据集**：有穷数据集合。

**方法集**：

- 初始化

- 判空

- 遍历（先序、中序、后序、层次）

```go
type BinaryTree[T cmp.Ordered] struct {
	Root *node[T]
}

type node[T cmp.Ordered] struct {
	value       T
	left, right *node[T]
}
```

详见 `/tree/binarytree/binarytree.go` 。

### 3.2.3 遍历

#### 深度优先

深度优先遍历优先访问距离根节点最远的节点。

对于任一节点，总是存在「当前节点」、「左子树」、「右子树」三部分，根据当前节点被访问的顺序，可以将二叉树的遍历分为**前**/**中**/**后**序遍历。

**前序遍历**：<u>当前节点</u> => 左子树 => 右子树。

```go
// 基于递归的前序遍历
func (t *BinaryTree[T]) TraverseInPreOrderByRecursion(vs *[]T) {
	t.traverseInPreOrderByRecursion(t.Root, vs)
}

func (t *BinaryTree[T]) traverseInPreOrderByRecursion(c *node[T], vs *[]T) {
	if c != nil {
		*vs = append(*vs, c.value)
		t.traverseInPreOrderByRecursion(c.left, vs)
		t.traverseInPreOrderByRecursion(c.right, vs)
	}
}
```

**中序遍历**：左子树 => <u>当前节点</u> => 右子树。

```go
// 基于递归的中序遍历
func (t *BinaryTree[T]) TraverseInInOrderByRecursion(vs *[]T) {
	t.traverseInInOrderByRecursion(t.Root, vs)
}

func (t *BinaryTree[T]) traverseInInOrderByRecursion(c *node[T], vs *[]T) {
	if c != nil {
		t.traverseInInOrderByRecursion(c.left, vs)
		*vs = append(*vs, c.value)
		t.traverseInInOrderByRecursion(c.right, vs)
	}
}
```

```go
// 基于堆栈的中序遍历
func (t *BinaryTree[T]) TraverseInInOrderByStack(vs *[]T) error {
	n, s := t.Root, &stack.LinkedListStack{}
	for {
		if n.left != nil {
			s.Push(n)
			n = n.left
			continue
		}
		*vs = append(*vs, n.value)
		if n.right != nil {
			n = n.right
			continue
		}
		if s.Empty() {
			break
		}
		p, e := s.Pop()
		if e != nil {
			return e
		}
		n = &node[T]{value: p.(*node[T]).value, right: p.(*node[T]).right}
	}
	return nil
}
```

**后序遍历**：左子树 => 右子树 => <u>当前节点</u>。

```go
// 基于递归的后序遍历
func (t *BinaryTree[T]) TraverseInPostOrderByRecursion(vs *[]T) {
	t.traverseInPostOrderByRecursion(t.Root, vs)
}

func (t *BinaryTree[T]) traverseInPostOrderByRecursion(c *node[T], vs *[]T) {
	if c != nil {
		t.traverseInPostOrderByRecursion(c.left, vs)
		t.traverseInPostOrderByRecursion(c.right, vs)
		*vs = append(*vs, c.value)
	}
}
```

```go
// 基于堆栈的后序遍历
func (t *BinaryTree[T]) TraverseInPostOrderByStack(vs *[]T) error {
	n, s := t.Root, &stack.LinkedListStack{}
	for {
		if n.left != nil {
			s.Push(&node[T]{value: n.value, right: n.right})
			n = n.left
			continue
		}
		if n.right != nil {
			r := n.right
			n.right = nil
			s.Push(n)
			n = r
			continue
		}
		*vs = append(*vs, n.value)
		if s.Empty() {
			break
		}
		p, e := s.Pop()
		if e != nil {
			return e
		}
		n = p.(*node[T])
	}
	return nil
}
```

详见 `/tree/binarytree/traverse_recursively.go` 和 `/tree/binarytree/traverse_by_stack.go` 。

#### 广度优先

广度优先遍历（也称为层序遍历）优先访问距离根节点最近且尚未被访问的节点。

```go
// 基于队列
func (t *BinaryTree[T]) TraverseBreadthFirstByQueue(vs *[]T) error {
	q := &queue.LinkedListQueue{}
	q.Put(t.Root)
	for !q.Empty() {
		a, e := q.Poll()
		if e != nil {
			return e
		}
		n := a.(*node[T])
		*vs = append(*vs, n.value)
		if n.left != nil {
			q.Put(n.left)
		}
		if n.right != nil {
			q.Put(n.right)
		}
	}
	return nil
}
```

```go
// 基于堆栈
func (t *BinaryTree[T]) TraverseBreadthFirstByStack(vs *[]T) error {
	current, next := &stack.LinkedListStack{}, &stack.LinkedListStack{}
	current.Push(t.Root)
	n := new(node[T])
	for !current.Empty() {
		if v, e := current.Pop(); e != nil {
			return e
		} else {
			n = v.(*node[T])
		}
		*vs = append(*vs, n.value)
		if n.right != nil {
			next.Push(n.right)
		}
		if n.left != nil {
			next.Push(n.left)
		}
		if current.Empty() {
			current, next = next, current
		}
	}
	return nil
}
```

详见 `/tree/binarytree/traverse_breadth_first.go` 。
