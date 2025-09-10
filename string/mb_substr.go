package phpstring

func MbSubstr(s string, start int, length ...int) string {
	runes := []rune(s)
	strLen := len(runes)

	if strLen == 0 {
		return ""
	}

	if start < 0 {
		start = strLen + start
		if start < 0 {
			start = 0
		}
	}

	if start >= strLen {
		return ""
	}

	if len(length) == 0 {
		return string(runes[start:])
	}

	l := length[0]
	if l == 0 {
		return ""
	}

	if l < 0 {
		endPos := strLen + l
		if endPos <= start {
			return ""
		}
		return string(runes[start:endPos])
	}

	endPos := start + l
	if endPos > strLen {
		endPos = strLen
	}

	return string(runes[start:endPos])
}
