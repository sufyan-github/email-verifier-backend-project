package main

import (
	"email-verifier/internal/handler"
	"email-verifier/pkg/config"
	"email-verifier/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	config.LoadConfig()

	// Init logger
	logger.InitLogger()

	r := gin.Default()

	r.GET("/verify", handler.VerifyHandler)

	port := config.AppConfig.AppPort

	logger.InfoLogger.Println("Server running on port:", port)

	r.Run(":" + port)
}