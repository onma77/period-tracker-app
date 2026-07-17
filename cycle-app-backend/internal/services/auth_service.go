package services

import (
	"errors"

	"cycle-app-backend/internal/models"
	"cycle-app-backend/internal/repository"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: repo,
	}
}


// type AuthService struct{}

// func NewAuthService() *AuthService {
// 	return &AuthService{}
// }

func (s *AuthService) Register(req models.RegisterRequest) (string, error) {

	if  req.FirstName == ""{
	return "", errors.New("first name is required")
	}

	if req.LastName == ""{
		return "", errors.New("last name is required")
	}

	if req.Email == ""{
		return "", errors.New("email is required")
	}
	
	if req.Password == ""{
		return "", errors.New("password is required")
	}

	if req.DateOfBirth == ""{
		return "", errors.New("date of birth required")
	}
	
	// Check if email already exists
	_, exists := s.userRepo.GetByEmail(req.Email)
	if exists {
		return "", errors.New("email already exists")
	}

	// Create user
	user := models.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PasswordHash: req.Password, // We'll hash this later
		DateOfBirth:  req.DateOfBirth,
	}

	s.userRepo.Create(user)
	// Later:
	// 1. Validate input
	// 2. Hash password
	// 3. Save user to database
	// 4. Send verification email

	return "Registration successful", nil
}