package l1_test

import (
	"math/rand"
	"testing"
)

// 变量和常量
func TestVarAndConst(t *testing.T) {
	const C = 1 // 声明一个常量，并赋值
	t.Log("C", C)
	var v1 int // 声明一个变量为int 类型，不赋值golang会先赋一个零值，这里是int类型，零值为0
	t.Log("v1 zero value", v1)
	v1 = 1 // 赋值变量
	t.Log("v1 = ", v1)
	v2 := 2 // 简短声明赋值，这种方式声明和赋值都一起
	t.Log("v2", v2)
	// 变量组
	var v3, v4 = 3, 4
	t.Log("v3", v3)
	t.Log("v4", v4)

	var (
		v5, v6 = 5, 6
		v7     = 7
		v8     = 8
	)
	t.Log("v5", v5)
	t.Log("v6", v6)
	t.Log("v7", v7)
	t.Log("v8", v8)

	v9, v10 := 9, 10
	t.Log("v9", v9)
	t.Log("v10", v10)

	// 什么时候用var 什么时候用 := ?
	// 一般情况下，声明的时候就有个初始值，都直接用 :=
	// 如果一开始没有初始值，后面会被重新赋值，或者需要指定类型，就用var
	isOpen := false // 默认关闭
	t.Log("isOpen", isOpen)
	var res int64 // 一开始就需要指定类型，但是先不赋值，在后面的判断中再赋值
	if rand.Int()%2 == 1 {
		res = 1
	} else {
		res = 0
	}
	t.Log("res", res)

}

// var a = b + c
// var b = aAddC()
// var c = 1

// func aAddC() int {
// 	return a + c
// }
// fmt
