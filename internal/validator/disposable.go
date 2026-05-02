// internal/validator/verifier.go
package validator

import "email-verifier/internal/models"

func VerifyEmail(email string) models.VerificationResult {
	result := models.VerificationResult{
		Email: email,
	}

	result.ValidSyntax = ValidateSyntax(email)

	if !result.ValidSyntax {
		return result
	}

	validDomain, mxRecords := ValidateDomain(email)
	result.ValidDomain = validDomain
	result.MXRecords = mxRecords

	if validDomain && len(mxRecords) > 0 {
		result.SMTPValid = VerifySMTP(email, mxRecords[0])
	}

	return result
}