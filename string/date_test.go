package phpstring

import (
	"strconv"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	fixed := time.Date(2025, time.September, 28, 15, 4, 5, 123456000, time.FixedZone("JST", 9*3600))

	originalNow := timeNow
	timeNow = func() time.Time { return fixed }
	defer func() { timeNow = originalNow }()

	tests := []struct {
		name   string
		layout string
		want   string
	}{
		{
			name:   "basic",
			layout: "Y-m-d",
			want:   "2025-09-28",
		},
		{
			name:   "escaped characters",
			layout: "\\T\\e\\s\\t Y",
			want:   "Test 2025",
		},
		{
			name:   "iso8601",
			layout: "c",
			want:   "2025-09-28T15:04:05+09:00",
		},
		{
			name:   "rfc2822",
			layout: "r",
			want:   "Sun, 28 Sep 2025 15:04:05 +0900",
		},
		{
			name:   "time formats",
			layout: "h:i:s A",
			want:   "03:04:05 PM",
		},
		{
			name:   "ordinal and month",
			layout: "jS F Y",
			want:   "28th September 2025",
		},
		{
			name:   "timezone",
			layout: "O P Z e",
			want:   "+0900 +09:00 32400 JST",
		},
		{
			name:   "unix timestamp",
			layout: "U",
			want:   strconv.FormatInt(fixed.Unix(), 10),
		},
		{
			name:   "empty layout",
			layout: "",
			want:   "",
		},
	}

	for _, tt := range tests {
		got := Date(tt.layout)
		if got != tt.want {
			t.Errorf("%s: expected %q, got %q", tt.name, tt.want, got)
		}
	}
}
