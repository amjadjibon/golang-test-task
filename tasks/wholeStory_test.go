package tasks

import (
	"testing"
)

func TestWholeStory(t *testing.T) {
	type input struct {
		s string
	}

	type output struct {
		story string
	}

	tests := []struct {
		input input
		want  output
	}{
		{input: input{s: "23-ab-48-caba-56-haha"}, want: output{story: "ab caba haha"}},
		{input: input{s: "1-hello-2-world"}, want: output{story: "hello world"}},
	}

	for _, test := range tests {
		got := wholeStory(test.input.s)
		if got != test.want.story {
			t.Errorf("wholeStory(%q) == %s, want %s", test.input.s, got, test.want.story)
		}
	}
}
