package sort

// 选择排序
func SelectionSort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[i] {
				Swap(list, i, j)
			}
		}
	}

}
