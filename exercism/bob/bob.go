package bob

import (
	"strings"
	"unicode"
)

func HasLetters(remark string) bool {

	flag := false

	for _, r := range remark {
		flag = flag || unicode.IsLetter(r)
	}

	return flag
}

// Makes Bob respond based on the remarks from Alice
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	question := strings.HasSuffix(remark, "?")
	yell := strings.Compare(remark, strings.ToUpper(remark)) == 0 && HasLetters(remark)
	empty := strings.Compare(strings.TrimSpace(remark), "") == 0

	switch {
	case question && yell:
		return "Calm down, I know what I'm doing!"
	case yell:
		return "Whoa, chill out!"
	case question:
		return "Sure."
	case empty:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

}
