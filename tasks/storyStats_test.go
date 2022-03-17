package tasks

import (
	"reflect"
	"testing"
)

func TestStoryStats(t *testing.T) {
	type input struct {
		s string
	}

	type output struct {
		shortest        string
		longest         string
		average         float64
		averageLenWords []string
	}

	tests := []struct {
		input input
		want  output
	}{
		{input: input{s: "1-hello-2-world"}, want: output{
			shortest:        "hello",
			longest:         "hello",
			average:         5.0,
			averageLenWords: []string{"hello", "world"},
		}},
	}

	for _, test := range tests {
		shortest, longest, average, averageLenWords := storyStats(test.input.s)
		if shortest != test.want.shortest {
			t.Errorf("storyState(%q) == %s, want %s", test.input.s, shortest, test.want.shortest)
		}

		if longest != test.want.longest {
			t.Errorf("storyState(%q) == %s, want %s", test.input.s, longest, test.want.longest)
		}

		if average != test.want.average {
			t.Errorf("storyState(%q) == %f, want %f", test.input.s, average, test.want.average)
		}

		if !reflect.DeepEqual(averageLenWords, test.want.averageLenWords) {
			t.Errorf("storyState(%q) == %s, want %s", test.input.s, averageLenWords, test.want.averageLenWords)
		}
	}
}
