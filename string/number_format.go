package phpstring

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

var msgPrinter = message.NewPrinter(language.Japanese)

func NumberFormat[N constraints.Integer | constraints.Float](n N) string {
	return msgPrinter.Sprint(number.Decimal(n))
}
