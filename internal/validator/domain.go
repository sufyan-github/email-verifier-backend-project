// internal/validator/domain.go
package validator

import (
	"net"
	"strings"
)

func ValidateDomain(email string) (bool, []string) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false, nil
	}

	domain := parts[1]

	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return false, nil
	}

	var mxHosts []string
	for _, mx := range mxRecords {
		mxHosts = append(mxHosts, mx.Host)
	}

	return true, mxHosts
}