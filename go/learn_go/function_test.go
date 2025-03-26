package learn_go

import (
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {
	function()
	v := function2()
	t.Run("有返回值的函数", func(t *testing.T) {
		if v != 1 {
			t.Errorf("v = %v", v)
		}
	})

	v1, v2 := function3()
	t.Run("多个返回值的函数", func(t *testing.T) {
		if v1 != 1 || v2 != 2 {
			t.Errorf("v1 = %v, v2 = %v", v1, v2)
		}
	})

	v3, v4 := function4()
	t.Run("命名的返回值", func(t *testing.T) {
		if v3 != 3 || v4 != 4 {
			t.Errorf("v3 = %v, v4 = %v", v3, v4)
		}
	})
}

// 无返回值的函数
func function() {
	fmt.Println("无返回值的函数")
}

//	func function2() int {
//		fmt.Println("有返回值的函数")
//		return 1
//	}
func function2() (v int) {
	fmt.Println("有返回值的函数")
	v = 1
	return
}

func function3() (int, int) {
	fmt.Println("有多个返回值的函数")
	return 1, 2
}

func function4() (v3 int, v4 int) {
	v3 = 3
	v4 = 4
	return
}
