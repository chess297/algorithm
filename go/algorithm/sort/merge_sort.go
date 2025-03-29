package sort

// 归并排序
func MergeSort(arr []int) {
	temp := make([]int, len(arr), cap(arr))
	copy(temp, arr)
	mergeSort(arr, 0, len(arr)-1, temp)
}

func mergeSort(arr []int, l, r int, temp []int) {
	if l >= r {
		return
	}
	m := l + (r-l)/2

	mergeSort(arr, l, m, temp)   // 递归左边数组，左边数组包含 m 位置，所以需要+1
	mergeSort(arr, m+1, r, temp) // 递归右边数组，右边数组不包含 m 位置，所以需要+1
	merge(arr, l, m, r, temp)
}

// 归并函数
func merge(arr []int, l, m, r int, temp []int) {
	copy(temp[l:r+1], arr[l:r+1])
	i := l
	j := m + 1
	for k := l; k <= r; k++ {
		if i > m { // 左边越界
			arr[k] = temp[j]
			j++
		} else if j > r { // 右边越界
			arr[k] = temp[i]
			i++
		} else if temp[i] < temp[j] { // 左边数小于右边数
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j] // 右边数小于或者等于左边数
			j++
		}
	}
}

// 通过协程实现归并排序
func MergeSortWithGoroutine() {

}
