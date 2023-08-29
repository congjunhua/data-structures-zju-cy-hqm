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

### 2.1.3 顺序存储

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

详见 `/linear_structure/linear_list/sequential.go` 。

### 2.1.4 链式存储

使用**链表**实现线性表。

数组实现的线性表，元素之间不仅逻辑上相邻，物理（存储介质）上也是相邻的，后者导致在插入和删除时，须手动更新受影响元素的物理位置。

而链表仅仅在逻辑上相邻，进而在插入和删除时不再需要移动受影响的元素，仅需修改链接即可。

```go
type node struct {
    value any
    next *node
}
```

其中：

- `next` 为指向下一个节点的指针。

只需要知道表头，即可访问任一位置上的元素。

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

使用数组实现堆栈，详见 `/linear_structure/stack/array.go` 和 `/linear_structure/stack/slice.go` 。

### 2.2.3 链式存储

使用链表实现堆栈，详见 `/linear_structure/stack/linked_list.go` 。

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

使用数组实现循环对列，详见 `/linear_structure/queue/array.go` 。

### 2.3.3 链式存储

使用链表实现对列，详见 `/linear_structure/queue/linked_list.go` 。

### 2.3.4 计算多项式

$p1=3x^5+4x^4-x^3+2x-1$

$p2=2x^4+x^3-7x^2+x$

求相加后的多项式 $p3$ 和相乘后的多项式 $p4$ 。

详见 `/linear_structure/queue/polynomial.go` 。线性结构

# 三、树

## 3.1 树与树的表示

### 3.1.1 顺序查找

即遍历数组，详见 `/tree/search/sequential.go` ，时间复杂度 $O(n)$ 。

### 3.1.2 二分查找

首先升序对数组进行排序，然后进入循环，每次循环，只对中间位置的元素与目标元素进行相等性和排序性比较：

1. 若中间位置的元素与目标元素相等，直接返回中间元素的下标；

2. 若中间位置的元素小于目标元素，说明目标元素必不可能出现在中间位置元素的左侧，待搜索区域缩小一半；

3. 若中间位置的元素大于目标元素，说明目标元素必不可能出现在中间位置元素的右侧，待搜索区域减小一半；

直至待搜索区域为零，详见 `/tree/search/binary.go` ，时间复杂度 $O(log_2\ n)$ 。

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

- **树的度**（Degree of tree）：所有节点的度的最大值。

- **叶节点**（Leaf）：度为零的节点。

- **层**（Level）：根节点为 $1$ 层，叶节点为最后一层，亦称为**深度**。

- **距离**（Distance）：两个节点间最短路径的边数。
