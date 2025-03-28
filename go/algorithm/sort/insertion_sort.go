package sort

// 插入排序
func InsertionSort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := i; j-1 >= 0; j-- {
			if list[j] < list[j-1] {
				Swap(list, j-1, j)
			} else {
				break
			}
		}
	}
}
