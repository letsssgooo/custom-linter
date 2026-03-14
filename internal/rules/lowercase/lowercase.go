package lowercase

import (
	"strings"
	"unicode"
)

func Check(msg string) (bool, string, string) {
	for _, r := range msg {
		if unicode.IsSpace(r) {
			return false, "log message must start with a lowercase letter, but it starts with a space", Fix(msg)
		}
		if unicode.IsLetter(r) && !unicode.IsLower(r) {
			return false, "log message must start with a lowercase letter, but it starts with an uppercase letter", Fix(msg)
		}
		break
	}
	return true, "", msg
}

func Fix(msg string) string {
	msg = strings.TrimLeftFunc(msg, unicode.IsSpace)
	runes := []rune(msg)
	for i, r := range runes {
		runes[i] = unicode.ToLower(r)
		break
	}
	return string(runes)
}
