// internal/service/verifier_service.go
package service

import (
	"email-verifier/internal/models"
	"email-verifier/internal/validator"
)

func VerifyEmail(email string) models.VerificationResult {
	result := models.VerificationResult{
		Email: email,
	}

	result.ValidSyntax = validator.ValidateSyntax(email)

	if !result.ValidSyntax {
		return result
	}

	result.IsDisposable = validator.IsDisposable(email)

	validDomain, mx := validator.ValidateDomain(email)
	result.ValidDomain = validDomain
	result.MXRecords = mx

	if validDomain && len(mx) > 0 {
		result.SMTPValid = validator.VerifySMTP(email, mx[0])
	}

	return result
}