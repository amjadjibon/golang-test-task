package tasks

import (
	"strings"
)

// Task: return the whole story
// Fist check if the story is already valid
// Time complexity: O(n)

func wholeStory(input string) string {
	var inputSlice = strings.Split(input, "-")
	var output = ""
	for i, v := range inputSlice {
		if i%2 != 0 {
			output += v
			output += " "
		}
	}

	return strings.TrimSpace(output)
}
