package phpstring

import (
	"strings"
	"testing"
)

func TestWordwrap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		opt   []WordwrapOpt
		want  string
	}{
		{
			input: "",
			want:  "",
		},
		{
			input: "aaaaa",
			want: "aaaaa",
		},
		{
			input: "abcdefghijklmnopqrstuvwxyz",
			opt: []WordwrapOpt{
				{
					Width: 7,
					Delim: "",
					CutLongWords: false,	
				},
			},
			want: strings.Join([]string{
				"abcdefg",
				"hijklmn",
				"opqrstu",
				"vwxyz",
			}, "\n"),
		},
		{
			input: "abcdefghijklmnopqrstuvwxyz",
			opt: []WordwrapOpt{
				{
					Width: 7,
					Delim: "",
					CutLongWords: true,	
				},
			},
			want: strings.Join([]string{
				"abcdefg",
				"hijklmn",
				"opqrstu",
				"vwxyz",
			}, "\n"),
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			Wordwrap(tt.input, tt.opt...)
		})
	}
}
