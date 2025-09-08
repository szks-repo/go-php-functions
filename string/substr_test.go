package phpstring

import "testing"

func TestSubstr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  string
		offset int
		length []int
		want   string
		desc   string
	}{
		{
			input:  "Hello World",
			offset: 0,
			length: nil,
			want:   "Hello World",
		},
		{
			input:  "Hello World",
			offset: 1,
			length: nil,
			want:   "ello World",
		},
		{
			input:  "Hello World",
			offset: 0,
			length: []int{5},
			want:   "Hello",
		},
		{
			input:  "Hello World",
			offset: 6,
			length: nil,
			want:   "World",
		},
		{
			input:  "Hello World",
			offset: -5,
			length: nil,
			want:   "World",
		},
		{
			input:  "Hello World",
			offset: -1,
			length: nil,
			want:   "d",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Substr(tt.input, tt.offset, tt.length...)
			if got != tt.want {
				t.Errorf("expected=%s, but got=%s", tt.want, got)
			}
		})
	}
}
