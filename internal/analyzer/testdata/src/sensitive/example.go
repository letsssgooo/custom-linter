package sensitive

import "log"

func example() {
	log.Print("token validated") // want "log message may contain sensitive data"
	log.Print("api_key-123")     // want "log message may contain sensitive data"
}
