package english

import "unicode"

func Check(msg string) (bool, string, string) {
	for _, r := range msg {
		if r > unicode.MaxASCII {
			return false, "log message must be in English", msg
		}
	}
	return true, "", msg
}
