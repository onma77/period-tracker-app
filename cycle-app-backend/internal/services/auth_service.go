package services

import "cycle-app-backend/internal/models"

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(req models.RegisterRequest) (string, error) {
	// Later:
	// 1. Validate input
	// 2. Hash password
	// 3. Save user to database
	// 4. Send verification email

	return "Registration successful", nil
}