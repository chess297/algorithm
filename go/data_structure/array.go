package data_structure

import "fmt"

type Array[T any] struct {
	// data [10]T // golang 中的数组，定义了长度就无法更改，所以只能用切片来处理
	data []T
	// 容量
	capacity int
	// 长度
	size int
}

func NewArray[T any](list []T, capacity int) Array[T] {
	return Array[T]{
		data:     list,
		capacity: capacity,
		size:     len(list),
	}
}

// 获取容量
func (a *Array[T]) GetCapacity() int {
	return a.capacity
}

func (a *Array[T]) GetSize() int {
	return len(a.data)
}

func (a *Array[T]) Add(e T, index int) {
	if index == 0 {
		a.data = append([]T{e}, a.data...)
	} else if index == a.GetSize() {
		a.data = append(a.data, e)
	} else {
		a.data = append(a.data[:index], append([]T{e}, a.data[index:]...)...)
	}
}

func (a *Array[T]) AddFirst(e T) {
	// a.data = append([]T{e}, a.data...)
	a.Add(e, 0)
}

func (a *Array[T]) AddLast(e T) {
	a.Add(e, a.GetSize())
}

func (a *Array[T]) Get(i int) (T, bool) {
	var zero T
	if i < 0 || i > a.GetSize() {
		return zero, false
	}
	return a.data[i], true
}

func (a *Array[T]) GetFirst() (T, bool) {
	return a.Get(0)
}

func (a *Array[T]) GetLast() (T, bool) {
	return a.Get(a.GetSize())
}

func (a *Array[T]) Delete(i int) {
	a.data = append(a.data[:i], a.data[i+1:]...)
}

func (a *Array[T]) DeleteLast() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	a.Delete(a.GetSize())
}

func (a *Array[T]) DeleteFirst() {
	a.Delete(0)
}
