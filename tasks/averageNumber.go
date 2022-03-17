package tasks

import (
	"strconv"
	"strings"
)

// Task: get average number from string
// First need to check validity of input string
// Then split string by space and get average number
// Time complexity: O(n)

func averageNumber(input string) float64 {
	var sum int
	var count int
	var inputSliced = strings.Split(input, "-")

	// Loop through the input string and add the numbers together
	for i, v := range inputSliced {
		// If the number is not a number, skip it
		if i%2 == 0 {
			// Convert the string to an int
			intVal, _ := strconv.Atoi(v)
			sum += intVal
			count++
		}
	}

	// Return the average
	return float64(sum) / float64(count)
}
