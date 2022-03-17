package tasks

import (
	"math"
	"strings"
)

// StoryStats is a struct that contains the stats for a story
// It is used to calculate the stats for a story
// First check if input is valid
// Time complexity: O(n)

func storyStats(input string) (shortest string, longest string, average float64, averageLenWords []string) {
	var shortsLen = 999999999
	var maxLen = 0
	var totalLen = 0

	var inputSliced = strings.Split(input, "-")
	for i, v := range inputSliced {
		if i%2 != 0 {
			if len(v) < shortsLen {
				shortsLen = len(v)
				shortest = v
			}

			if len(v) > maxLen {
				maxLen = len(v)
				longest = v
			}

			totalLen += len(v)
		}
	}

	average = float64(totalLen) / float64(len(inputSliced)/2)

	for i, v := range inputSliced {
		if i%2 != 0 {
			if len(v) == int(math.Ceil(average)) || len(v) == int(math.Floor(average)) {
				averageLenWords = append(averageLenWords, v)
			}
		}
	}

	return shortest, longest, average, averageLenWords
}
