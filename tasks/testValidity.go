package tasks

import (
	"strconv"
	"strings"
)

func testValidity(input string) bool {
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

		}
	}

	return true
}
