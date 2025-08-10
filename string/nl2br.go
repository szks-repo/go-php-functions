package phpstring

import "strings"

var nl2brReplacer = strings.NewReplacer("\n", "<br />")

func Nl2br(s string) string {
	return nl2brReplacer.Replace(s)
}
