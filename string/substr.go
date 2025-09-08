package phpstring

func Substr(s string, offset int, length ...int) string {
	// Convert string to rune slice to handle UTF-8 properly
	runes := []rune(s)
	strLen := len(runes)

	// Handle empty string
	if strLen == 0 {
		return ""
	}

	// Handle negative offset (count from end)
	if offset < 0 {
		offset = strLen + offset
		if offset < 0 {
			offset = 0
		}
	}

	// If offset is beyond string length, return empty string
	if offset >= strLen {
		return ""
	}

	// Determine the length to extract
	var extractLen int
	if len(length) > 0 {
		extractLen = length[0]

		// Handle negative length (exclude characters from end)
		if extractLen < 0 {
			extractLen = strLen - offset + extractLen
			if extractLen < 0 {
				return ""
			}
		}
	} else {
		// If no length specified, extract from offset to end
		extractLen = strLen - offset
	}

	// Calculate end position
	end := offset + extractLen
	if end > strLen {
		end = strLen
	}

	// Extract substring and convert back to string
	if offset < end {
		return string(runes[offset:end])
	}

	return ""
}
