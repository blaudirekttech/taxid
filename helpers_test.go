package taxid

import (
	"math"
	"reflect"
	"testing"
)

func TestConvertToint64Slice(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedSlice  []int64
		expectedResult bool
	}{
		{
			name:           "Empty string",
			input:          "",
			expectedSlice:  nil,
			expectedResult: false,
		},
		{
			name:           "Valid string",
			input:          "1234",
			expectedSlice:  []int64{1, 2, 3, 4},
			expectedResult: true,
		},
		{
			name:           "String containing leading zeros",
			input:          "00123",
			expectedSlice:  []int64{0, 0, 1, 2, 3},
			expectedResult: true,
		},
		{
			name:           "String containing non-digit characters",
			input:          "12a3",
			expectedSlice:  nil,
			expectedResult: false,
		},
		{
			name:           "String containing leading whitespace",
			input:          " 123",
			expectedSlice:  []int64{1, 2, 3},
			expectedResult: true,
		},
		{
			name:           "String containing trailing whitespace",
			input:          "123 ",
			expectedSlice:  []int64{1, 2, 3},
			expectedResult: true,
		},
		{
			name:           "String containing both leading and trailing whitespace",
			input:          " 123 ",
			expectedSlice:  []int64{1, 2, 3},
			expectedResult: true,
		},
		{
			name:           "String containing only whitespace",
			input:          "   ",
			expectedSlice:  nil,
			expectedResult: false,
		},
		{
			name:           "String containing a single digit",
			input:          "4",
			expectedSlice:  []int64{4},
			expectedResult: true,
		},
		{
			name:           "String containing a large number",
			input:          "1234567890",
			expectedSlice:  []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			expectedResult: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			slice, result := convertToInt64Slice(tc.input)
			if result != tc.expectedResult {
				t.Errorf("Expected result %v, but got %v", tc.expectedResult, result)
			}
			if !reflect.DeepEqual(slice, tc.expectedSlice) {
				t.Errorf("Expected slice %v, but got %v", tc.expectedSlice, slice)
			}
		})
	}
}

func TestUniqueCharsInAString(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput int64
	}{
		{
			name:           "Empty string",
			input:          "",
			expectedOutput: 0,
		},
		{
			name:           "String with all unique characters",
			input:          "abcdefg",
			expectedOutput: 7,
		},
		{
			name:           "String with repeated characters",
			input:          "hello world",
			expectedOutput: 8,
		},
		{
			name:           "String with only one character",
			input:          "a",
			expectedOutput: 1,
		},
		{
			name:           "String with spaces and non-alphabetic characters",
			input:          "!@#$%^&*()     zyxwvutsrqponmlkjihgfedcba",
			expectedOutput: 37,
		},
		{
			name:           "String with all the same character",
			input:          "aaaaaaaaaa",
			expectedOutput: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := uniqueCharsInAString(tc.input)
			if output != tc.expectedOutput {
				t.Errorf("Expected %d unique characters, but got %d", tc.expectedOutput, output)
			}
		})
	}
}

func TestSumDigits(t *testing.T) {
	testCases := []struct {
		name           string
		input          int64
		expectedOutput int64
	}{
		{
			name:           "Zero",
			input:          0,
			expectedOutput: 0,
		},
		{
			name:           "Single digit number",
			input:          5,
			expectedOutput: 5,
		},
		{
			name:           "Number with multiple digits",
			input:          123456,
			expectedOutput: 21,
		},
		{
			name:           "Negative number",
			input:          -123,
			expectedOutput: 6,
		},
		{
			name:           "Number with trailing zeros",
			input:          10000,
			expectedOutput: 1,
		},
		{
			name:           "Number with max value",
			input:          math.MaxInt64,
			expectedOutput: 88,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := sumDigits(tc.input)
			if output != tc.expectedOutput {
				t.Errorf("Expected %d, but got %d", tc.expectedOutput, output)
			}
		})
	}
}
