// syntax.go
package validator

import "net/mail"

func ValidateSyntax(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}