package learn_go

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestString(t *testing.T) {
	// 字符串
	str := "hello world"
	t.Run("字符串的长度", func(t *testing.T) {
		if len(str) != 11 {
			t.Errorf("len(str) = %v, want %v", len(str), 11)
		}
	})

	str = "你好，世界"
	t.Run("中文字符串的长度", func(t *testing.T) {
		if utf8.RuneCountInString(str) != 5 {
			t.Errorf("utf8.RuneCountInString(str), want %v", 5)
		}
	})

	str1 := str + "~"
	t.Run("字符串的拼接", func(t *testing.T) {
		if str1 != "你好，世界~" {
			t.Errorf("str1 = %v, want %v", str1, "你好，世界~")
		}
	})

	str2 := fmt.Sprintf("%s, %s", str, "golang")

	t.Run("字符串的格式化", func(t *testing.T) {
		if str2 != "你好，世界, golang" {
			t.Errorf("str2 = %v, want %v", str2, "你好，世界, golang")
		}
	})

	str3 := str2[0:12]
	t.Run("字符串的截取", func(t *testing.T) {
		if str3 != "你好，世" { // 一个中文占3个字节
			t.Errorf("str3 = %v, want %v", str3, "你好，世")
		}
	})
}
