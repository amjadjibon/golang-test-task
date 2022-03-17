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
		{input: input{s: "23-ab-48-caba-56-haha"}, want: output{average: (23 + 48 + 56) / 3.0}},
		{input: input{s: "1-hello-2-world"}, want: output{average: (1 + 2) / 2.0}},
	}

	for _, test := range tests {
		got := averageNumber(test.input.s)
		if got != test.want.average {
			t.Errorf("averageNumber(%q) == %f, want %f", test.input.s, got, test.want.average)
		}
	}
}
