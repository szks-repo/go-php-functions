package string

import "strings"

type WordwrapOpt struct {
	Width        int
	Delim        string
	CutLongWords bool
}

func Wordwrap(s string, opt ...WordwrapOpt) string {
	if s == "" {
		return s
	}

	var width = 75
	var delim = "\n"
	var cutLong bool

	if len(opt) == 1 {
		if opt[0].Width > 0 {
			width = opt[0].Width
		}
		if opt[0].Delim != "" {
			delim = opt[0].Delim
		}
		cutLong = opt[0].CutLongWords
	}

	lines := strings.Split(s, "\n")
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		result = append(result, wrapLine(line, width, delim, cutLong))
	}

	return strings.Join(result, "\n")
}

func wrapLine(line string, width int, delim string, cutLong bool) string {
	if cutLong {
		return wrapLineNoCut(line, width, delim)
	} else {
		return wrapLineCut(line, width, delim)
	}
}

func wrapLineNoCut(str string, width int, delim string) string {
	words := splitIntoWordBytes(str)
	if len(words) == 0 {
		return ""
	}

	var res strings.Builder
	var currentLineLen int
	var needSpace bool

	for _, word := range words {
		if word == " " || word == "\t" {
			if currentLineLen > 0 && currentLineLen < width {
				res.WriteString(word)
				currentLineLen++
			}
			needSpace = false
		} else {
			wordLen := len(word)
			if needSpace && currentLineLen > 0 {
				if currentLineLen+1+wordLen > width {
					res.WriteString(delim)
					currentLineLen = 0
					needSpace = false
				} else {
					res.WriteString("")
					currentLineLen++
				}
			}

			if currentLineLen == 0 || currentLineLen+wordLen <= width {
				res.WriteString(word)
				currentLineLen += wordLen
			} else {
				res.WriteString(delim)
				res.WriteString(word)
				currentLineLen = wordLen
			}
			needSpace = true
		}
	}

	return res.String()
}

func wrapLineCut(str string, width int, delim string) string {
	if str == "" {
		return ""
	}

	var res strings.Builder

	var strAsBytes = []byte(str)
	var lineLen int
	var i int
	for i < len(strAsBytes) {
		for i < len(strAsBytes) && (strAsBytes[i] == ' ' || strAsBytes[i] == '\t') {
			i++
		}
		if i >= len(strAsBytes) {
			break
		}

		wordStart := i
		for i < len(strAsBytes) && strAsBytes[i] != ' ' && strAsBytes[i] != '\t' {
			i++
		}

		wordLen := i - wordStart
		if wordLen == 0 {
			if lineLen < width {
				res.WriteByte(strAsBytes[wordStart])
				lineLen++
			}
			i++
			continue
		}

		if lineLen == 0 {
			// start of line
			if wordLen > width {
				res.Write(strAsBytes[wordStart:width])
				res.WriteString(delim)
				lineLen = 0
			} else {
				res.Write(strAsBytes[wordStart:i])
				lineLen = wordLen
			}
		} else {
			// middle of line
			if lineLen+wordLen > width {
				if wordLen > width {
					res.WriteString(delim)
					lineLen = 0
					i = wordStart
				} else {
					res.WriteString(delim)
					res.Write(strAsBytes[wordStart:i])
					lineLen = wordLen
				}
			} else {
				res.WriteString(" ")
				res.Write(strAsBytes[wordStart:i])
				lineLen += 1 + wordLen
			}
		}
	}

	return res.String()
}

func splitIntoWordBytes(s string) []string {
	var words []string
	var currentWord strings.Builder
	for _, c := range s {
		if c == ' ' || c == '\t' {
			if currentWord.Len() > 0 {
				words = append(words, string(c))
				currentWord.Reset()
			}
			words = append(words, string(c))
		} else {
			currentWord.WriteRune(c)
		}
	}
	// flush
	if currentWord.Len() > 0 {
		words = append(words, currentWord.String())
	}

	return words
}
