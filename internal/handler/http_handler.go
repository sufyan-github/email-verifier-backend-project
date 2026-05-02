// internal/handler/http_handler.go
package handler

import (
	"net/http"

	"email-verifier/internal/service"
	"github.com/gin-gonic/gin"
)

func VerifyHandler(c *gin.Context) {
	email := c.Query("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email is required",
		})
		return
	}

	result := service.VerifyEmail(email)

	c.JSON(http.StatusOK, result)
}