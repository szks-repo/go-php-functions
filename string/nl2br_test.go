package phpstring

import "testing"

func TestNl2br(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  string
	}{
		{
			input: "Hello\nWorld",
			want:  "Hello<br />World",
		},
		{
			input: "Line1\nLine2\nLine3",
			want:  "Line1<br />Line2<br />Line3",
		},
		{
			input: "NoNewline",
			want:  "NoNewline",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := Nl2br(tt.input)
			if got != tt.want {
				t.Errorf("Nl2br(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
