package concepts

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"time"
)

/*
	讨论1.3 再试一个多项式
	https://www.icourse163.org/learn/ZJU-93001?tid=1470659487#/learn/content?type=detail&id=1254705090&cid=1286093695
*/

// Polynomial 分别使用普通算法和秦九韶算法计算多项式，并对比时间复杂度。
func Polynomial(n int, x float64) {
	cs := make([]float64, n+1)
	cs[0] = 1
	for i := 1; i < len(cs); i++ {
		cs[i] = 1 / float64(i)
	}
	for _, f := range []func([]float64, float64) float64{common, qinJiuShao} {
		t := time.Now()
		fmt.Println(f(cs, x))
		fmt.Printf(
			"%s耗时：%v\n",
			runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(),
			time.Now().Sub(t),
		)
	}
}

func common(cs []float64, x float64) float64 {
	p := float64(0)
	for i := 0; i < len(cs); i++ {
		p += cs[i] * math.Pow(x, float64(i))
	}
	return p
}

func qinJiuShao(cs []float64, x float64) float64 {
	p := cs[len(cs)-1]
	for i := len(cs) - 2; i >= 0; i-- {
		p = cs[i] + x*p
	}
	return p
}
