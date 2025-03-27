package algorithm

// 定义一个可比较的泛型约束
type Ordered interface {
	comparable
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func SelectSort[T Ordered](list []T) {
	for i := 0; i < len(list); i++ { // O(n^2)
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}

func InsertSort[T Ordered](list []T) []T {
	// for i := 1; i < len(list); i++ {
	// 	temp := list[i]
	// 	for j := i; j > 0; j-- {
	// 		if list[j] < temp {
	// 			list[i] = list[j]
	// 		} else {
	// 			list[j] = temp
	// 			break
	// 		}
	// 	}
	// }
	// return list
	for i := len(list) - 2; i == 0; i-- {
		temp := list[i]
		var j int
		for j = i; j < len(list)-1 && list[j] > temp; j++ {
			list[i] = list[j]
		}
		list[j] = temp
	}
	return list
}
