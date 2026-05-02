// internal/models/result.go
package models

type VerificationResult struct {
	Email        string   `json:"email"`
	ValidSyntax  bool     `json:"valid_syntax"`
	ValidDomain  bool     `json:"valid_domain"`
	MXRecords    []string `json:"mx_records"`
	SMTPValid    bool     `json:"smtp_valid"`
	IsDisposable bool     `json:"is_disposable"`
}