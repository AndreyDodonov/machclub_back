package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/AndreyDodonov/machclub_back/pkg/models"
	"github.com/AndreyDodonov/machclub_back/pkg/repository"
)

const salt = "nekiy_nabor_sluchaynyh_simvolov_1"

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
func (s *AuthService) generatePasswordHash (password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}