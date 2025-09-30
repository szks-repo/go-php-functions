package phpstring

import (
	"strconv"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	fixed := time.Date(2025, time.September, 28, 15, 4, 5, 123456000, time.FixedZone("JST", 9*3600))
	leap := time.Date(2024, time.February, 29, 6, 30, 45, 0, time.FixedZone("JST", 9*3600))
	isoBoundary := time.Date(2022, time.January, 1, 8, 0, 0, 0, time.UTC)

	originalNow := timeNow
	timeNow = func() time.Time { return fixed }
	defer func() { timeNow = originalNow }()

	tests := []struct {
		name   string
		layout string
		want   string
		now    *time.Time
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
			name:   "hour variants",
			layout: "H-G-g-h a A",
			want:   "15-15-3-03 pm PM",
		},
		{
			name:   "ordinal and month",
			layout: "jS F Y",
			want:   "28th September 2025",
		},
		{
			name:   "month variants",
			layout: "y/n M D l N w",
			want:   "25/9 Sep Sun Sunday 7 0",
		},
		{
			name:   "day of year and iso week",
			layout: "z o W",
			want:   "270 2025 39",
		},
		{
			name:   "timezone",
			layout: "T O P Z e",
			want:   "JST +0900 +09:00 32400 JST",
		},
		{
			name:   "microseconds and milliseconds",
			layout: "u v",
			want:   "123456 123",
		},
		{
			name:   "days in month and leap",
			layout: "t L",
			want:   "30 0",
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
		{
			name:   "leap year days",
			layout: "t L",
			want:   "29 1",
			now:    &leap,
		},
		{
			name:   "iso week boundary",
			layout: "o-W-N",
			want:   "2021-52-6",
			now:    &isoBoundary,
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			target := fixed
			if test.now != nil {
				target = *test.now
			}

			timeNow = func() time.Time { return target }
			t.Cleanup(func() {
				timeNow = func() time.Time { return fixed }
			})

			got := Date(test.layout)
			if got != test.want {
				t.Errorf("layout %q: expected %q, got %q", test.layout, test.want, got)
			}
		})
	}
}
