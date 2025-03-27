package data_structure

type Array[T any] struct {
	data []T
}

func (a *Array[T]) GetSize() int {
	return len(a.data)
}

func (a *Array[T]) Add(e T) {
	a.data = append(a.data, e)
}

func (a *Array[T]) Find(i int) *T {
	if i > len(a.data) {
		return nil
	}
	return &a.data[i]
}

func (a *Array[T]) Delete(i int) {
	if i < 0 || i >= a.GetSize() {
		panic(" 不存在指定元素 ")
	}
	a.data = append(a.data[:i], a.data[i+1:]...)
}
