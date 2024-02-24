package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var loginRequest LoginRequest
	err := c.Bind(&loginRequest)
	if err != nil {
		return c.JSON(400, err)
	}
	loginResponse, err := h.authService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return echo.ErrUnauthorized.WithInternal(err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = loginResponse.Token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.JSON(200, loginResponse)
}
