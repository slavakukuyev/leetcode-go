package validparentheses

import "testing"

func BenchmarkIsValidParentheses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidParentheses("(()[]{{}})")
	}
}

func BenchmarkIsValidParenthesesCopilot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidParenthesesCopilot("(()[]{{}})")
	}
}

func TestIsValidParentheses(t *testing.T) {
	if !IsValidParentheses("(()[]{{}})") {
		t.Error("Expected true, got false")
	}
}

func TestIsValidParentheses2(t *testing.T) {
	if !IsValidParentheses("()") {
		t.Error("Expected true, got false")
	}
}

func TestIsValidParentheses3(t *testing.T) {
	if IsValidParentheses(")(") {
		t.Error("Expected false, got true")
	}
}

func TestIsValidParentheses4(t *testing.T) {
	if IsValidParentheses("(]") {
		t.Error("Expected false, got true")
	}
}

func TestIsValidParenthesesCopilot(t *testing.T) {
	if !IsValidParenthesesCopilot("(()[]{{}})") {
		t.Error("Expected true, got false")
	}
}

func TestIsValidParenthesesCopilot2(t *testing.T) {
	if !IsValidParenthesesCopilot("()") {
		t.Error("Expected true, got false")
	}
}

func TestIsValidParenthesesCopilot3(t *testing.T) {
	if IsValidParenthesesCopilot(")(") {
		t.Error("Expected false, got true")
	}
}

func TestIsValidParenthesesCopilot4(t *testing.T) {
	if IsValidParenthesesCopilot("(]") {
		t.Error("Expected false, got true")
	}
}

func TestIsValidParenthesesCopilot5(t *testing.T) {
	if IsValidParenthesesCopilot("([)]") {
		t.Error("Expected false, got true")
	}
}
