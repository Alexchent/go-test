package util

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strconv"
)

func PrintK(number int) (withCommaThousandSep string) {
	p := message.NewPrinter(language.Chinese)
	withCommaThousandSep = p.Sprintf("%f", number)
	return
}

func AddCommas(n int64) string {
	s := strconv.FormatInt(n, 10)
	result := ""
	for i, c := range s {
		if i != 0 && (len(s)-i)%3 == 0 {
			result += ","
		}
		result += string(c)
	}
	return result
}
