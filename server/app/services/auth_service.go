// app/services/auth_service.go
package services

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/omelab/go-crm/app/repositories"
)

type AuthService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}
 

func (s *AuthService) Login(email string, password string) (string, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return "", errors.New("Invalid email or password")
		}
		return "", err
	}

	err = s.userRepository.ComparePasswords(user.Password, password)
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	// Add any other claims you want

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", errors.New("Failed to generate token")
	}

	return tokenString, nil
}