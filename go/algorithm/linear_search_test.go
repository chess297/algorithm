package algorithm

import "testing"

type _Case[T comparable] struct {
	name   string
	list   []T
	target T
	want   int
}

func TestLinearSearchOnInt(t *testing.T) {
	tests := []_Case[int]{
		{
			name:   "found",
			list:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "not found",
			list:   []int{1, 2, 3, 4, 5},
			target: 6,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinearSearch(tt.list, tt.target); got != tt.want {
				t.Errorf("LinearSearch() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestLinearSearchOnString(t *testing.T) {
	tests := []_Case[string]{
		{
			name:   "not found",
			list:   []string{"1", "2", "3", "4", "5"},
			target: "6",
			want:   -1,
		},
		{
			name:   "found",
			list:   []string{"1", "2", "3", "4", "5"},
			target: "1",
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinearSearch(tt.list, tt.target); got != tt.want {
				t.Errorf("LinearSearch() = %v, want %v", got, tt.want)
			}
		})
	}

}
