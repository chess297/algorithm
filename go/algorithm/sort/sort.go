package sort

// 将数组内i,j两个位置交换
func Swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
