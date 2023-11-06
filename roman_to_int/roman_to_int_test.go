package romantoint

// The function returns 4 when passed "IV".
import "testing"

func TestRomanToIntWithIV(t *testing.T) {
	result := RomanToInt("IV")
	if result != 4 {
		t.Errorf("Expected 4, but got %d", result)
	}
}

// The function returns 9 when passed "IX".
func TestRomanToIntWithIX(t *testing.T) {
	result := RomanToInt("IX")
	if result != 9 {
		t.Errorf("Expected 9, but got %d", result)
	}
}

func TestRomanToIntWithEmptyString(t *testing.T) {
	result := RomanToInt("")
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}

func TestRomanToIntWithNonRomanCharacters(t *testing.T) {
	result := RomanToInt("ABC")
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}

func TestRomanToIntWithInvalidRomanNumeral(t *testing.T) {
	result := RomanToInt("IIII")
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}

//create a few test functions to validate the RomanToInt function
// Path: roman_to_int/roman_to_int.go
func TestRomanToIntWithValidRomanNumeral(t *testing.T) {
	result := RomanToInt("III")
	if result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}
}

func TestRomanToIntWithValidRomanNumeral2(t *testing.T) {
	result := RomanToInt("XII")
	if result != 12 {
		t.Errorf("Expected 12, but got %d", result)
	}
}

func TestRomanToIntWithValidRomanNumeral3(t *testing.T) {
	result := RomanToInt("XVI")
	if result != 16 {
		t.Errorf("Expected 16, but got %d", result)
	}
}

func TestRomanToIntWithValidRomanNumeral4(t *testing.T) {
	result := RomanToInt("XXVII")
	if result != 27 {
		t.Errorf("Expected 27, but got %d", result)
	}
}

func TestRomanToIntWithValidRomanNumeral5(t *testing.T) {
	result := RomanToInt("XLIV")
	if result != 44 {
		t.Errorf("Expected 44, but got %d", result)
	}
}

func TestRomanToIntWithValidRomanNumeral6(t *testing.T) {
	result := RomanToInt("LXVIII")
	if result != 68 {
		t.Errorf("Expected 68, but got %d", result)
	}
}

func TestRomanToIntWithValidRomanNumeral7(t *testing.T) {
	result := RomanToInt("LXXXIII")
	if result != 83 {
		t.Errorf("Expected 83, but got %d", result)
	}
}

func TestRomanToIntWithValidRomanNumeral8(t *testing.T) {
	result := RomanToInt("XCIX")
	if result != 99 {
		t.Errorf("Expected 99, but got %d", result)
	}
}

func TestRomanToIntWithValidRomanNumeral9(t *testing.T) {
	result := RomanToInt("DCCCXC")
	if result != 890 {
		t.Errorf("Expected 890, but got %d", result)
	}
}

func TestRomanToIntWithValidRomanNumeral10(t *testing.T) {
	result := RomanToInt("MCMXCIV")
	if result != 1994 {
		t.Errorf("Expected 1994, but got %d", result)
	}
}

func TestRomanToIntWithValidRomanNumeral11(t *testing.T) {
	result := RomanToInt("MMCDXX")
	if result != 2420 {
		t.Errorf("Expected 2420, but got %d", result)
	}
}

func TestRomanToIntWithValidRomanNumeral12(t *testing.T) {
	result := RomanToInt("MMMCMXCIX")
	if result != 3999 {
		t.Errorf("Expected 3999, but got %d", result)
	}
}

func TestValidRomanLength1(t *testing.T) {
	result := isValidRoman("I")
	if result != true {
		t.Errorf("Expected true, but got false")
	}
}

// Returns true for valid roman numerals with length > 1.
func TestValidRomanLengthGreaterThan1(t *testing.T) {
	result := isValidRoman("IV")
	if result != true {
		t.Errorf("Expected true, but got false")
	}
}

// Returns true for valid roman numerals with all valid characters.
func TestValidRomanAllValidCharacters(t *testing.T) {
	result := isValidRoman("XVIII")
	if result != true {
		t.Errorf("Expected true, but got false")
	}
}

// Returns false for roman numerals with length > 15.
func TestInvalidRomanLengthGreaterThan15(t *testing.T) {
	result := isValidRoman("MMMMMMMMMMMMMMM")
	if result != false {
		t.Errorf("Expected false, but got true")
	}
}

func TestInvalidRomanInvalidCharacters(t *testing.T) {
	result := isValidRoman("IXL")
	if result != false {
		t.Errorf("Expected false, but got true")
	}
}
