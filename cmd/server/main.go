// cmd/main.go
package main

import (
	"encoding/json"
	"fmt"
	"email-verifier/internal/validator"
)

func main() {
	email := "test@example.com"

	result := validator.VerifyEmail(email)

	jsonResult, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonResult))
}