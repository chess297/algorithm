// channel 是用于协程之间通信和传输数据的
package learn_go

import (
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	t.Run("channel ", func(t *testing.T) {
		ch := make(chan int)
		go func() {
			ch <- 1
		}()
		go func() {
			data := <-ch
			if data != 1 {
				t.Errorf(" data is %v want %d", data, 1)
			}
		}()
		time.Sleep(time.Second * 1)
	})
}
