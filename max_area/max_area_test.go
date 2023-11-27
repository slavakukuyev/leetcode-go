package maxarea

import "testing"

func BenchmarkMaxArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	}
}

func BenchmarkMaxAreaCopilot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxAreaCopilot([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	}
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "test1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "test2",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "test3",
			height: []int{4, 3, 2, 1, 4},
			want:   16,
		},
		{
			name:   "test4",
			height: []int{1, 2, 1},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxArea(tt.height); got != tt.want {
				t.Errorf("MaxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
