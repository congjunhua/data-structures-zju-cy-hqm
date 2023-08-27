package concepts

/*
	最大连续子列和（最大子数组和）问题。
	https://leetcode.cn/problems/maximum-subarray/
*/

func MaxSubArray(nums []int) int {
	return onlineProcess(nums)
}

// 【算法一】暴力计算所有可能的连续部分
func normal(nums []int) int {
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

// 【算法二】减少重复的加法运算，以消除第三轮循环
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

// 【算法三】分治法
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

// 【算法四】在线处理
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
