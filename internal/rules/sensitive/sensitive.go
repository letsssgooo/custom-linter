package sensitive

import "strings"

var data = []string{
	"password",
	"passwd",
	"pwd",
	"token",
	"api_key",
	"apikey",
	"secret",
	"private_key",
	"auth",
	"authorization",
	"bearer",
	"session",
	"cookie",
}

func Check(msg string) (bool, string, string) {
	lower := strings.ToLower(msg)
	for _, k := range data {
		if strings.Contains(lower, k) {
			return false, "log message may contain sensitive data", msg
		}
	}
	return true, "", msg
}
