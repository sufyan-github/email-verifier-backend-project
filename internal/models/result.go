// internal/models/result.go
package models

type VerificationResult struct {
	Email        string
	ValidSyntax  bool
	ValidDomain  bool
	MXRecords    []string
	SMTPValid    bool
}