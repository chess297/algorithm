package algorithm

// params：
// list: 待查找的列表
// target: 待查找的目标值
// returns:
// int: 目标值在列表中的索引，如果不存在则返回-1
func LinearSearch(list []int, target int) int {
	for i, v := range list {
		if v == target {
			// found
			return i
		}
	}
	// not found
	return -1
}
