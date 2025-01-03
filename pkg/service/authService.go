package service

import (
	"os"
	"time"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/models"
	"github.com/MDmitryM/banking-app-go/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

type jwtBankingClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user bankingApp.User) (string, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	var userModel = models.UserModel{
		Email:         user.Email,
		Username:      user.Username,
		Password_hash: string(pwdHash),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	return s.repo.CreateUser(userModel)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	id, err := s.repo.IsUserValid(email, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtBankingClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	signedJWTToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return signedJWTToken, nil
}
