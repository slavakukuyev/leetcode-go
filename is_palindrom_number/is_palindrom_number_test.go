package ispalindromnumber

import "testing"

func BenchmarkIsPalindromNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindromeNumber(12345678987654321)
	}
}

func BenchmarkIsPalindromCopilot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindromNumberCopilot(12345678987654321)
	}
}

func TestIsPalindromNumber(t *testing.T) {

	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"1", 12345678987654321, true},
		{"2", 12345678987654322, false},
		{"3", 12345678987654323, false},
		{"4", 12345678987654324, false},
		{"5", 12345678987654325, false},
		{"6", 12345678987654326, false},
		{"7", 12345678987654327, false},
		{"8", 12345678987654328, false},
		{"9", 12345678987654329, false},
		{"10", 12345678987654320, false},
		{"11", -12345678987654321, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindromeNumber(tt.x); got != tt.want {
				t.Errorf("IsPalindromNumber(%d) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}

func TestIsPalindromNumberCopilot(t *testing.T) {

	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"1", 12345678987654321, true},
		{"2", 12345678987654322, false},
		{"3", 12345678987654323, false},
		{"4", 12345678987654324, false},
		{"5", 12345678987654325, false},
		{"6", 12345678987654326, false},
		{"7", 12345678987654327, false},
		{"8", 12345678987654328, false},
		{"9", 12345678987654329, false},
		{"10", 12345678987654320, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindromNumberCopilot(tt.x); got != tt.want {
				t.Errorf("IsPalindromNumberCopilot(%d) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}
