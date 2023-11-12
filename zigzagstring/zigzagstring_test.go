package zigzagstring

import "testing"

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert("PAYPALISHIRING", 3)
	}
}

func BenchmarkConvertCopilot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertCopilot("PAYPALISHIRING", 3)
	}
}

func TestConvertReturnsInputStringIfNumRowsIsGreaterThanOrEqualToLengthOfString(t *testing.T) {
	input := "hello"
	numRows := 6
	expected := "hello"

	result := Convert(input, numRows)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Returns the input string if numRows is 1 or the length of the string is 1
func TestConvertReturnsInputStringIfNumRowsIs1(t *testing.T) {
	input := "hello"
	numRows := 1
	expected := "hello"

	result := Convert(input, numRows)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Returns the expected output for a valid input string and numRows greater than 1
func TestConvertReturnsExpectedOutputForValidInputAndNumRowsGreaterThan1(t *testing.T) {
	input := "PAYPALISHIRING"
	numRows := 3
	expected := "PAHNAPLSIIGYIR"

	result := Convert(input, numRows)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Returns the expected output for a valid input string and numRows equal to the length of the string
func TestConvertReturnsExpectedOutputForValidInputAndNumRowsEqualToLengthOfString(t *testing.T) {
	input := "hello"
	numRows := 5
	expected := "hello"

	result := Convert(input, numRows)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Returns an empty string if the input string is empty
func TestConvertReturnsEmptyStringIfInputStringIsEmpty(t *testing.T) {
	input := ""
	numRows := 2
	expected := ""

	result := Convert(input, numRows)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Returns the input string if numRows is less than or equal to 0
func TestConvertReturnsInputStringIfNumRowsIsLessThanOrEqualTo0(t *testing.T) {
	input := "hello"
	numRows := 0
	expected := "hello"

	result := Convert(input, numRows)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestConvert(t *testing.T) {
	input := "PAYPALISHIRING"
	numRows := 4
	expected := "PINALSIGYAHRPI"

	result := Convert(input, numRows)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestConvertCopilot(t *testing.T) {
	input := "PAYPALISHIRING"
	numRows := 4
	expected := "PINALSIGYAHRPI"

	result := ConvertCopilot(input, numRows)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
