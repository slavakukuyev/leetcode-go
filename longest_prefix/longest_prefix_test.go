package longestprefix

import "testing"

func BenchmarkLongestPrefix(b *testing.B) {
	strs := []string{"flowerskjaghajdfhgaksjhfgkajfgka", "flowfgjdfjdsfjdfhdkh", "flight", "fler", "flo", "flp", "fl", "flituret", "flsadasd"}
	for i := 0; i < b.N; i++ {
		LongestCommonPrefix(strs)
	}
}

func BenchmarkLongestPrefixLeetCode(b *testing.B) {
	strs := []string{"flowerskjaghajdfhgaksjhfgkajfgka", "flowfgjdfjdsfjdfhdkh", "flight", "fler", "flo", "flp", "fl", "flituret", "flsadasd"}
	for i := 0; i < b.N; i++ {
		LongestCommonPrefixLeetCode(strs)
	}
}

func BenchmarkLongestCommonPrefixCopilot(b *testing.B) {
	strs := []string{"flowerskjaghajdfhgaksjhfgkajfgka", "flowfgjdfjdsfjdfhdkh", "flight", "fler", "flo", "flp", "fl", "flituret", "flsadasd"}
	for i := 0; i < b.N; i++ {
		LongestCommonPrefixCopilot(strs)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	expected := "fl"
	actual := LongestCommonPrefix(strs)
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestLongestCommonPrefixLeeCode(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	expected := "fl"
	actual := LongestCommonPrefixLeetCode(strs)
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestLongestCommonPrefixCopilot(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	expected := "fl"
	actual := LongestCommonPrefixCopilot(strs)
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

//create failing test
func TestLongestCommonPrefixEmptyResult(t *testing.T) {
	strs := []string{"", "flow", "flight"}
	expected := ""
	actual := LongestCommonPrefix(strs)
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestLongestCommonPrefixLeeCodeEmptyResult(t *testing.T) {
	strs := []string{"", "flow", "flight"}
	expected := ""
	actual := LongestCommonPrefixLeetCode(strs)
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestLongestCommonPrefixCopilotEmptyResult(t *testing.T) {
	strs := []string{"", "flow", "flight"}
	expected := ""
	actual := LongestCommonPrefixCopilot(strs)
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestLongestCommonPrefixEmptyInput(t *testing.T) {
	strs := []string{}
	expected := ""
	actual := LongestCommonPrefix(strs)
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestLongestCommonPrefixArrayWithEmptyStrings(t *testing.T) {
	strs := []string{"", "", ""}
	expected := ""
	actual := LongestCommonPrefix(strs)
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestLongestCommonPrefixArrayWithDifferentStrings(t *testing.T) {
	strs := []string{"flower", "flow", "flight", "cardfg"}
	expected := ""
	actual := LongestCommonPrefix(strs)
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

// Returns the only string in the list when the list has only one string.
func TestSingleString(t *testing.T) {
	strs := []string{"hello"}
	expected := "hello"
	result := LongestCommonPrefix(strs)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Returns an empty string when any string in the list is longer than 200 characters.
func TestLongString(t *testing.T) {
	strs := []string{"flower", "flow", "a very long string that is longer than 200 characters.........................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................."}
	expected := ""
	result := LongestCommonPrefix(strs)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Returns an empty string when the input list has more than 200 strings.
func TestManyStrings(t *testing.T) {
	strs := make([]string, 201)
	for i := 0; i < 201; i++ {
		strs[i] = "string"
	}
	expected := ""
	result := LongestCommonPrefix(strs)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Returns an empty string when the first string in the list is longer than 200 characters.
func TestLongFirstString(t *testing.T) {
	strs := []string{"fl a very long string that is longer than 200 characters..........................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................", "flow", "flight"}
	expected := ""
	result := LongestCommonPrefix(strs)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
