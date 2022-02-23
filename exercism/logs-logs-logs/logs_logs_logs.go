package logs

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {

	var application string

	for _, c := range log {
		char := fmt.Sprintf("%c", c)

		if application == "" || application == "default" {
			switch char {
			case "❗":
				application = "recommendation"
			case "🔍":
				application = "search"
			case "☀":
				application = "weather"
			default:
				application = "default"
			}
		}

	}

	return application
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runeLog := []rune(log)
	for idx, rune := range runeLog {
		if rune == oldRune {
			runeLog[idx] = newRune
		}
	}
	return string(runeLog)

}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
