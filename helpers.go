package taxid

import (
	"strconv"
	"strings"
)

func convertToInt64Slice(str string) ([]int64, bool) {
	str = strings.ReplaceAll(str, " ", "")
	if _, err := strconv.Atoi(str); err != nil {
		return nil, false
	}

	var slice []int64

	for _, digit := range str {
		slice = append(slice, int64(digit)-int64('0'))
	}

	return slice, true
}

func uniqueCharsInAString(str string) int64 {
	count := int64(0)

	for i := 0; i < len(str); i++ {
		appears := false

		for j := 0; j < i; j++ {
			if str[j] == str[i] {
				appears = true

				break
			}
		}

		if !appears {
			count++
		}
	}

	return count
}

func sumDigits(number int64) int64 {
	if number < 0 {
		number = -number
	}

	sumResult := int64(0)
	for number > 0 {
		sumResult += number % 10
		number /= 10
	}

	return sumResult
}
