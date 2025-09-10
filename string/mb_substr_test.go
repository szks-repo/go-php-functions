package phpstring

import "testing"

func TestMbSubstr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s        string
		start    int
		length   []int
		expected string
	}{
		{"Hello, 世界", 0, nil, "Hello, 世界"},
		{"Hello, 世界", 7, nil, "世界"},
		{"Hello, 世界", -2, nil, "世界"},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			result := MbSubstr(test.s, test.start, test.length...)
			if result != test.expected {
				t.Errorf("MbSubstr(%q, %d, %v) = %q; want %q", test.s, test.start, test.length, result, test.expected)
			}
		})
	}
}
