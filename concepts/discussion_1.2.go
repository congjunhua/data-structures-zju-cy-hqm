package concepts

/*
	讨论1.2 晒一下PrintN在你的机器上运行的结果？
	https://www.icourse163.org/learn/ZJU-93001?tid=1470659487#/learn/content?type=detail&id=1254705090&cid=1286093693
*/

// PrintN 依次使用循环和递归的方式，顺序打印从 1 到 N 的所有数字。
func PrintN(n uint) {
	for _, f := range []func(uint){loop, recur} {
		f(n)
	}
}

func loop(n uint) {
	for i := 1; i <= int(n); i++ {
		println(i)
	}
}

func recur(n uint) {
	if n > 0 {
		recur(n - 1)
		println(n)
	}
}
