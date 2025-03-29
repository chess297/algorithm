// 并发
// 为什么需要并发？
// 1. 防止死锁
// 2. IO操作
// 3. 扩展性
package learn_go

import (
	"sync"
	"testing"
	"time"
)

var (
	students = map[int64]*student{
		1: {
			id:   1,
			name: "s1",
			age:  13,
		},
		2: {
			id:   2,
			name: "s2",
			age:  13,
		},
		3: {
			id:   3,
			name: "s3",
			age:  13,
		},
		4: {
			id:   4,
			name: "s4",
			age:  13,
		},
	}
)

type student struct {
	id   int64
	name string
	age  int64
}

func getStudent(i int64) *student {
	time.Sleep(time.Millisecond * 200)
	return students[i]
}

func TestMain(t *testing.T) {
	t.Run("TestMain goroutine ", func(t *testing.T) {
		var wg sync.WaitGroup
		var mutex sync.Mutex
		ids := []int64{1, 2}
		result := map[int64]int64{}
		slice := []int64{}
		// for _, id := range ids { // for range 是无序的，所以id 可能会有问题
		for i := 0; i <= len(ids)-1; i++ {
			wg.Add(1)
			id := ids[i]
			go func(id int64) {
				defer func() {
					mutex.Unlock()
					wg.Done()
				}()
				mutex.Lock()
				s := getStudent(id)
				result[id] = s.age
				slice = append(slice, id)
			}(id)
		}
		wg.Wait()
		if result[1] != 13 {
			t.Errorf("result[1] is %d want 15", result[1])
		}
		// if slice[0] != 1 {
		// 	t.Errorf("slice[0] is %d want 1", slice[0])
		// }
		// if slice[1] != 2 {
		// 	t.Errorf("slice[1] is %d want 2", slice[1])
		// }
	})

}
