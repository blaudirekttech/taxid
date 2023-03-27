package taxid

import (
	"math"
	"strconv"
)

func convertToIntSlice(str string) ([]int, bool) {
	if _, err := strconv.Atoi(str); err != nil {
		return nil, false
	}

	var slice []int

	for _, digit := range str {
		slice = append(slice, int(digit)-int('0'))
	}

	return slice, true
}

func uniqueCharsInAString(str string) int {
	count := 0

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

func sumDigits(number int) int {
	remainder := 0
	sumResult := 0
	for number != 0 {
		remainder = number % 10
		sumResult += remainder
		number = number / 10
	}
	return sumResult
}

func digitAt(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
