package l1_test

import "testing"

func TestNil(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		var a interface{}
		var i *int64
		t.Log(a == nil)
		t.Log(i == nil)
	})

}
