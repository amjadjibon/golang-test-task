package tasks

import (
	"strconv"
	"strings"
	"unicode"
)

func isASCII(s string) bool {
	// ASCII is a subset of UTF-8
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}

	return true
}

// TestValidity tests if the given string is valid.
//
func testValidity(input string) bool {
	// Check if the input is empty
	if len(input) == 0 {
		return false
	}

	var inputSliced = strings.Split(input, "-")

	if len(inputSliced)%2 != 0 {
		return false
	}

	for i, v := range inputSliced {
		if i%2 == 0 {
			intVal, err := strconv.Atoi(v)
			if err != nil {
				return false
			}

			if intVal < 0 {
				return false
			}
		} else {
			if len(v) == 0 {
				return false
			}

			if !isASCII(v) {
				return false
			}
		}
	}

	return true
}
