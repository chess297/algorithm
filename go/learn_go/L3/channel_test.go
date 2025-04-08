// channel 是用于协程之间通信和传输数据的
package learn_go

import (
	"sync"
	"testing"
)

func TestChannel(t *testing.T) {
	t.Run("channel 的基础使用", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)
		ch := make(chan int)
		go func() {
			defer wg.Done()
			ch <- 1
		}()
		go func() {
			defer wg.Done()
			data := <-ch
			if data != 1 {
				t.Errorf(" data is %v want %d", data, 1)
			}
		}()
		wg.Wait()
	})
}
