package service

import (
	"errors"
	"reast-api/internal/models"
	"reast-api/internal/repository"

	"golang.org/x/crypto/bcrypt"
	// Import your preferred JWT package here
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(req models.RegisterRequest) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Prepare data for BOTH tables
	newUser := models.UserCredential{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Profile: models.UserProfile{
			FullName: req.FullName,
		},
	}

	return s.repo.CreateUser(&newUser)
}

func (s *AuthService) Login(req models.LoginRequest) (string, error) {
	// 1. Find user by email OR username
	user, err := s.repo.FindByIdentifier(req.Identifier)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// 2. Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// 3. Generate JWT Token (Pseudo-code, implement your JWT logic here)
	// token := GenerateJWT(user.ID, user.Username)
	token := "dummy-jwt-token-replace-me"

	return token, nil
}
