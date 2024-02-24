package auth

import (
	"github.com/labstack/gommon/log"
	"hrm-api/src/internal/jwt"
	"hrm-api/src/internal/user"
	"hrm-api/src/internal/utils"
)

type AuthService interface {
	Login(username string, password string) (LoginResponse, error)
}

type authService struct {
	userService user.UserService
}

func NewAuthService(memberService user.UserService) AuthService {
	return &authService{memberService}
}

func (a authService) Login(username string, password string) (LoginResponse, error) {
	m, err := a.userService.GetByUsername(username)
	if err != nil {
		log.Errorf("Error logging in: %s", err)
		return LoginResponse{}, err
	}

	err = utils.CompareHashAndPassword(m.Password, password)
	if err != nil {
		log.Errorf("Error logging in: %s", err)
		return LoginResponse{}, err
	}

	token := jwt.GenerateToken(m.Username, m.IsAdmin)

	if err != nil {
		log.Errorf("Error logging in: %s", err)
		return LoginResponse{}, err
	}

	return LoginResponse{token, m.ID, m.IsAdmin}, nil
}
