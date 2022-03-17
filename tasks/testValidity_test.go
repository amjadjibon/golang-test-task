package tasks

import (
	"testing"
)

func TestTestValidity(t *testing.T) {
	type input struct {
		s string
	}

	type output struct {
		valid bool
	}

	tests := []struct {
		input input
		want  output
	}{
		{input: input{s: "23-ab-48-caba-56-haha"}, want: output{valid: true}},
		{input: input{s: "1-hello-2-world"}, want: output{valid: true}},
		{input: input{s: ""}, want: output{valid: false}},
		{input: input{s: "a-b"}, want: output{valid: false}},
		{input: input{s: "hello world"}, want: output{valid: false}},
	}

	for _, test := range tests {
		got := testValidity(test.input.s)
		if got != test.want.valid {
			t.Errorf("testValidity(%q) == %t, want %t", test.input.s, got, test.want.valid)
		}
	}
}
