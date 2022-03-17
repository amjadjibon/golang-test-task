package tasks

import (
	"testing"
)

func TestAverageNumber(t *testing.T) {
	type input struct {
		s string
	}

	type output struct {
		average float64
	}

	tests := []struct {
		input input
		want  output
	}{
		{input: input{s: "23-ab-48-caba-56-haha"}, want: output{average: 42.3333333333}},
		{input: input{s: "1-hello-2-world"}, want: output{average: 1.5}},
	}

	for _, test := range tests {
		got := averageNumber(test.input.s)
		if got != test.want.average {
			t.Errorf("averageNumber(%q) == %f, want %f", test.input.s, got, test.want.average)
		}
	}
}
