package search

// params：
// list: 待查找的列表
// target: 待查找的目标值
// returns:
// int: 目标值在列表中的索引，如果不存在则返回-1
func LinearSearch[T comparable](list []T, target T) int {
	for i, v := range list {
		if v == target {
			return i
		}
	}
	return -1
}
