package algorithm

import "testing"

type _Test struct {
	name   string
	list   []int
	target int
	want   int
}

func TestLinearSearch(t *testing.T) {
	tests := []_Test{}
	tests = append(tests, _Test{
		name:   "found",
		list:   []int{1, 2, 3, 4, 5},
		target: 3,
		want:   2,
	})
	tests = append(tests, _Test{
		name:   "not found",
		list:   []int{1, 2, 3, 4, 5},
		target: 6,
		want:   -1,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinearSearch(tt.list, tt.target); got != tt.want {
				t.Errorf("LinearSearch() = %v, want %v", got, tt.want)
			}
		})
	}

}
