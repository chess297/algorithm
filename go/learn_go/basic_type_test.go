// Golang 中的基础类型
// 1. 布尔类型
// 2. 数字类型

package learn_go

import "testing"

func TestBasicType(t *testing.T) {
	// 1. 布尔类型
	var b bool = true
	t.Log(b)
	// 2. 数字类型
	// 有符号整数
	var i int = 1
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	// 无符号整数
	var ui uint = 1
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615
	// 浮点数
	var f32 float32 = 3.1415926         // 精确到7位
	var f64 float64 = 3.141592653589793 // 精确到15位
	// 复数
	var c complex64 = 1 + 2i    // 复数，实部和虚部都是float32类型
	var c64 complex128 = 1 + 2i // 复数，实部和虚部都是float64类型
	t.Log(i, i8, i16, i32, i64, ui, ui8, ui16, ui32, ui64, f32, f64, c, c64)

}
