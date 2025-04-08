package l1_test

import (
	"testing"
)

func TestMap(t *testing.T) {
	var m map[string]int
	// m["one"] = 1                 // 会报错，因为 map 是一个引用类型，需要初始化
	{
		m = make(map[string]int, 10) // 初始化 map，10 是容量，不是长度
	}
	m["one"] = 2

	t.Run("map 长度", func(t *testing.T) {
		if len(m) != 1 {
			t.Errorf("len(m) = %v, want %v", len(m), 1)
		}
	})

	for k, v := range m {
		t.Log(k, v)
	}

}
