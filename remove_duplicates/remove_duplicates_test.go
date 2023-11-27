package removeduplicates

import "testing"

func BenchmarkRemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveDuplicates([]int{1, 1, 1, 2, 2, 3})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test1",
			nums: []int{1, 1, 1, 2, 2, 3},
			want: 5,
		},
		{
			name: "test2",
			nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want: 7,
		},
		{
			name: "test3",
			nums: []int{1, 1, 1},
			want: 2,
		},
		{
			name: "test4",
			nums: []int{1, 1, 1, 1, 2, 2, 3},
			want: 5,
		},
		{
			name: "test5",
			nums: []int{1, 1, 1, 1, 2, 2, 2, 3},
			want: 5,
		},
		{
			name: "test6",
			nums: []int{1, 1, 1, 1, 2, 2, 2, 3, 3},
			want: 6,
		},
		{
			name: "test7",
			nums: []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3},
			want: 6,
		},
		{
			name: "test8",
			nums: []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3},
			want: 6,
		},
		{
			name: "test9",
			nums: []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4},
			want: 7,
		},
		{
			name: "test10",
			nums: []int{0, 0, 0, 0, 0},
			want: 2,
		},
		{
			name: "test11",
			nums: []int{1},
			want: 1,
		},
		{
			name: "test12",
			nums: []int{1, 1},
			want: 2,
		},
		{
			name: "test13",
			nums: []int{1, 1, 1},
			want: 2,
		},
		{
			name: "test14",
			nums: []int{1, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicates(tt.nums); got != tt.want {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
