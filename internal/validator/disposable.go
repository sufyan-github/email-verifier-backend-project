// disposable.go
package validator

import "strings"

var disposableDomains = []string{
	"mailinator.com",
	"10minutemail.com",
	"tempmail.com",
}

func IsDisposable(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	domain := parts[1]

	for _, d := range disposableDomains {
		if domain == d {
			return true
		}
	}

	return false
}