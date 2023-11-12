package zigzagstring

import "bytes"

func Convert(s string, numRows int) string {
	l := len(s)
	if numRows < 2 || l < 2 {
		return s
	}

	currentIndex := 0
	stack := [][]byte{}
	ascending := true
	tempArray := make([]byte, numRows)
	for i := 0; i < l; i++ {

		if ascending {
			tempArray[currentIndex] = s[i]
			currentIndex++
		} else {
			if currentIndex != 0 && currentIndex != numRows-1 {
				tempArray[currentIndex] = s[i]
			} else {
				i--
			}
			currentIndex--
		}

		if currentIndex == numRows || currentIndex == -1 {
			stack = append(stack, tempArray)
			ascending = !ascending
			tempArray = make([]byte, numRows)
		}

		if currentIndex == numRows {
			currentIndex--
		} else if currentIndex == -1 {
			currentIndex++
		}
	}

	stack = append(stack, tempArray)

	var str bytes.Buffer
	ls := len(stack)
	for i := 0; i < numRows; i++ {
		for j := 0; j < ls; j++ {
			if stack[j][i] != 0 {
				str.WriteByte(stack[j][i])
			}
		}
	}

	return str.String()
}

func ConvertCopilot(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var result []byte
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += 2*numRows - 2 {
			result = append(result, s[j])
			if i != 0 && i != numRows-1 && j+2*numRows-2-2*i < len(s) {
				result = append(result, s[j+2*numRows-2-2*i])
			}
		}
	}
	return string(result)
}
