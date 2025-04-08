// 并发
// 为什么需要并发？
// 1. 防止死锁
// 2. IO操作
// 3. 扩展性
package l2_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestGoroutine(t *testing.T) {

	t.Run("协程的执行顺序是不确定的", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			fmt.Println(1)
		}()
		go func() {
			defer wg.Done()
			fmt.Println(2)

		}()
		wg.Wait()
		fmt.Println("执行完成")
	})

}

func TestMutex(t *testing.T) {

	t.Run("通过互斥锁来确保执行顺序", func(t *testing.T) {
		var wg sync.WaitGroup
		var mu sync.Mutex
		turn := 1
		wg.Add(2)
		go func() {
			defer wg.Done()
			for {
				mu.Lock()
				if turn == 1 {
					turn = 2
					fmt.Println(turn)
					mu.Unlock()
					break
				}
				mu.Unlock()
			}

		}()
		go func() {
			defer wg.Done()
			for {
				mu.Lock()

				if turn == 2 {
					turn = 3
					fmt.Println(turn)
					mu.Unlock()
					break
				}
				mu.Unlock()
			}

		}()
		wg.Wait()
		fmt.Println("执行完成")
	})

}
