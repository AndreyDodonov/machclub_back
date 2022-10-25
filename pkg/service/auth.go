package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/AndreyDodonov/machclub_back/pkg/models"
	"github.com/AndreyDodonov/machclub_back/pkg/repository"

	"github.com/golang-jwt/jwt/v4"
)

const (
	salt       = "nekiy_nabor_sluchaynyh_simvolov_1" // TODO remove to .env
	tokenTTL   = 12 * time.Hour
	signingKey = "nekiy_klyuch_podpisi_3"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// create new user and hashing his password
func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// password hash generation
func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, s.generatePasswordHash(password))
	if err != nil {
		return "", nil
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(), // TODO remove to config
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	t, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return t, nil
}
