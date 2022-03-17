package tasks

import (
	"strconv"
	"strings"
)

func averageNumber(input string) float64 {
	var sum int
	var count int
	var inputSliced = strings.Split(input, "-")

	for i, v := range inputSliced {
		if i%2 == 0 {
			intVal, _ := strconv.Atoi(v)
			sum += intVal
			count++
		}
	}

	return float64(sum) / float64(count)
}
