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
