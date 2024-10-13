package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pysk0101/todo-app-mux/backend/internal/core/domain"
	"github.com/pysk0101/todo-app-mux/backend/internal/core/ports"
)

type AuthServiceImpl struct {
	userRepo ports.UserRepository
}

func NewAuthServiceImpl(userRepo ports.UserRepository) ports.AuthService {
	return &AuthServiceImpl{userRepo: userRepo}
}

func (a *AuthServiceImpl) Register(user *domain.User) error {
	// Burada kullanıcı kaydı için gerekli işlemler yapılır (örneğin, şifreyi hash'leme)
	return a.userRepo.Create(user)
}

func (a *AuthServiceImpl) Login(email, password string) (string, error) {
	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil || user.Password != password { // Şifre karşılaştırma yapılmalı
		return "", errors.New("invalid username or password")
	}

	token, err := a.generateToken(int(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *AuthServiceImpl) generateToken(userID int) (string, error) {

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET is NULL")
	}

	expirationTime := time.Now().Add(30 * time.Minute) // Token geçerlilik süresi
	claims := &jwt.StandardClaims{
		Subject:   fmt.Sprint(userID),
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
