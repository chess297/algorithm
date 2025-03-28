package array

// 切片数组
type SliceArray struct {
	// golang 中的数组，编译的时候就确定了类型，所以只能用切片
	data []int
}

// 创建一个数组
func NewSliceArray(list []int) SliceArray {
	return SliceArray{
		data: list,
	}
}

// 获取容量
func (a *SliceArray) GetCapacity() int {
	return cap(a.data)
}

// 获取数组长度
func (a *SliceArray) GetSize() int {
	return len(a.data)
}

// 获取数组长度
func (a *SliceArray) IsEmpty() bool {
	return a.GetSize() == 0
}

// 在指定位置添加元素
func (a *SliceArray) Add(e int, index int) {
	if index == 0 {
		a.data = append([]int{e}, a.data...)
	} else if index == a.GetSize() {
		a.data = append(a.data, e)
	} else {
		// a.data[:index] = [0:index]
		a.data = append(a.data[:index], append([]int{e}, a.data[index:]...)...)
	}
}

// 在首位添加元素
func (a *SliceArray) AddFirst(e int) {
	a.Add(e, 0)
}

// 在最后添加元素
func (a *SliceArray) AddLast(e int) {
	a.Add(e, a.GetSize())
}

// 获取元素
func (a *SliceArray) Get(i int) int {
	return a.data[i]
}

// 获取数组第一个元素
func (a *SliceArray) GetFirst() int {
	return a.Get(0)
}

// 获取数组最后一个元素
func (a *SliceArray) GetLast() int {
	return a.Get(a.GetSize() - 1)
}

// 删除元素
func (a *SliceArray) Delete(i int) {
	a.data = append(a.data[:i], a.data[i+1:]...)
}

// 删除最后一个元素
func (a *SliceArray) DeleteLast() {
	a.Delete(a.GetSize() - 1)
}

// 删除最后一个元素
func (a *SliceArray) DeleteFirst() {
	a.Delete(0)
}
