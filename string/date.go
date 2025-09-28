package phpstring

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var timeNow = time.Now

func Date(layout string) string {
	return formatWith(layout, timeNow())
}

func formatWith(layout string, t time.Time) string {
	if layout == "" {
		return ""
	}

	var b strings.Builder
	runes := []rune(layout)
	for i := 0; i < len(runes); i++ {
		ch := runes[i]
		if ch == '\\' {
			if i+1 < len(runes) {
				b.WriteRune(runes[i+1])
				i++
			}
			continue
		}

		switch ch {
		case 'Y':
			b.WriteString(fmt.Sprintf("%04d", t.Year()))
		case 'y':
			b.WriteString(fmt.Sprintf("%02d", t.Year()%100))
		case 'm':
			b.WriteString(fmt.Sprintf("%02d", int(t.Month())))
		case 'n':
			b.WriteString(strconv.Itoa(int(t.Month())))
		case 'M':
			b.WriteString(t.Format("Jan"))
		case 'F':
			b.WriteString(t.Format("January"))
		case 'd':
			b.WriteString(fmt.Sprintf("%02d", t.Day()))
		case 'j':
			b.WriteString(strconv.Itoa(t.Day()))
		case 'D':
			b.WriteString(t.Format("Mon"))
		case 'l':
			b.WriteString(t.Format("Monday"))
		case 'N':
			b.WriteString(strconv.Itoa(isoWeekday(t)))
		case 'w':
			b.WriteString(strconv.Itoa(int(t.Weekday())))
		case 'z':
			b.WriteString(strconv.Itoa(t.YearDay() - 1))
		case 'S':
			b.WriteString(ordinalSuffix(t.Day()))
		case 't':
			b.WriteString(strconv.Itoa(daysInMonth(t)))
		case 'L':
			b.WriteString(boolTo01(isLeapYear(t.Year())))
		case 'o':
			year, _ := t.ISOWeek()
			b.WriteString(fmt.Sprintf("%04d", year))
		case 'W':
			_, week := t.ISOWeek()
			b.WriteString(fmt.Sprintf("%02d", week))
		case 'H':
			b.WriteString(fmt.Sprintf("%02d", t.Hour()))
		case 'G':
			b.WriteString(strconv.Itoa(t.Hour()))
		case 'h':
			b.WriteString(fmt.Sprintf("%02d", hour12(t)))
		case 'g':
			b.WriteString(strconv.Itoa(hour12(t)))
		case 'i':
			b.WriteString(fmt.Sprintf("%02d", t.Minute()))
		case 's':
			b.WriteString(fmt.Sprintf("%02d", t.Second()))
		case 'u':
			b.WriteString(fmt.Sprintf("%06d", t.Nanosecond()/1000))
		case 'v':
			b.WriteString(fmt.Sprintf("%03d", t.Nanosecond()/1e6))
		case 'a':
			if t.Hour() < 12 {
				b.WriteString("am")
			} else {
				b.WriteString("pm")
			}
		case 'A':
			if t.Hour() < 12 {
				b.WriteString("AM")
			} else {
				b.WriteString("PM")
			}
		case 'T':
			b.WriteString(t.Format("MST"))
		case 'O':
			b.WriteString(t.Format("-0700"))
		case 'P':
			b.WriteString(t.Format("-07:00"))
		case 'Z':
			_, offset := t.Zone()
			b.WriteString(strconv.Itoa(offset))
		case 'U':
			b.WriteString(strconv.FormatInt(t.Unix(), 10))
		case 'e':
			b.WriteString(t.Location().String())
		case 'c':
			b.WriteString(formatWith("Y-m-d\\TH:i:sP", t))
		case 'r':
			b.WriteString(formatWith("D, d M Y H:i:s O", t))
		default:
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func isoWeekday(t time.Time) int {
	wd := int(t.Weekday())
	if wd == 0 {
		return 7
	}
	return wd
}

func ordinalSuffix(day int) string {
	if day%10 == 1 && day%100 != 11 {
		return "st"
	}
	if day%10 == 2 && day%100 != 12 {
		return "nd"
	}
	if day%10 == 3 && day%100 != 13 {
		return "rd"
	}
	return "th"
}

func daysInMonth(t time.Time) int {
	firstOfNextMonth := time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, t.Location())
	return firstOfNextMonth.Day()
}

func boolTo01(v bool) string {
	if v {
		return "1"
	}
	return "0"
}

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	return year%4 == 0
}

func hour12(t time.Time) int {
	hour := t.Hour() % 12
	if hour == 0 {
		return 12
	}
	return hour
}
