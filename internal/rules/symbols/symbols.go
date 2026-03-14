package symbols

import "unicode"

func Check(msg string) (bool, string, string) {
	for _, r := range msg {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			continue
		}
		switch r {
		case '-', '_':
			continue
		default:
			return false, "log message must not contain special symbols or emoji", Fix(msg)
		}
	}
	return true, "", msg
}

func Fix(msg string) string {
	out := make([]rune, 0, len(msg))
	for _, r := range msg {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) || r == '-' || r == '_' {
			out = append(out, r)
		}
	}
	return string(out)
}
