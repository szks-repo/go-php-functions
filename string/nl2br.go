package phpstring

import "strings"

var nl2brReplacer = strings.NewReplacer("\n", "<br />")

func Nl2br[S ~string](s S) string {
	return nl2brReplacer.Replace(string(s))
}
