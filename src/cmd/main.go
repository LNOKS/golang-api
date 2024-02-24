package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hrm-api/src/internal/auth"
	"hrm-api/src/internal/database"
	member "hrm-api/src/internal/user"
	"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	dbClient := database.Initialize()

	memberService := member.NewUserService(member.NewUserRepository(dbClient.DB))
	authService := auth.NewAuthService(memberService)
	authHandler := auth.NewAuthHandler(authService)

	setupAuthRoutes(e, authHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func setupAuthRoutes(e *echo.Echo, authHandler *auth.AuthHandler) {
	v1 := e.Group("/api/v1")
	authGroup := v1.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
}
