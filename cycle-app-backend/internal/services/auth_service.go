package services

import (
	"cycle-app-backend/internal/models"
	"cycle-app-backend/internal/repository"
	"cycle-app-backend/internal/utils"
	"errors"
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

	if req.FirstName == "" {
		return "", errors.New("first name is required")
	}

	if req.LastName == "" {
		return "", errors.New("last name is required")
	}

	if req.Email == "" {
		return "", errors.New("email is required")
	}

	if req.Password == "" {
		return "", errors.New("password is required")
	}

	if req.DateOfBirth == "" {
		return "", errors.New("date of birth required")
	}

	// Check if email already exists
	_, exists := s.userRepo.GetByEmail(req.Email)
	if exists {
		return "", errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	user := models.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		DateOfBirth:  req.DateOfBirth,
	}

	// Create user
	// user := models.User{
	// 	FirstName:    req.FirstName,
	// 	LastName:     req.LastName,
	// 	Email:        req.Email,
	// 	PasswordHash: req.Password, // We'll hash this later
	// 	DateOfBirth:  req.DateOfBirth,
	// }

	s.userRepo.Create(user)
	// Later:
	// 1. Validate input
	// 2. Hash password
	// 3. Save user to database
	// 4. Send verification email

	return "Registration successful", nil
}

func (s *AuthService) Login(req models.LoginRequest) (string, string, error) {

	if req.Email == "" {
		return "", "", errors.New("email is required")
	}

	if req.Password == "" {
		return "", "", errors.New("password is required")
	}

	user, exists := s.userRepo.GetByEmail(req.Email)

	if !exists {
		return "", "", errors.New("invalid email or password")
	}

	err := utils.CheckPassword(user.PasswordHash, req.Password)

	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		return "", "", errors.New("could not generate token")
	}

	return "Login successful", token, nil
}

// func (s *AuthService) Login(req models.LoginRequest) (string, error) {

// 	if req.Email == "" {
// 		return "", errors.New("email is required")
// 	}

// 	if req.Password == "" {
// 		return "", errors.New("password is required")
// 	}

// 	user, exists := s.userRepo.GetByEmail(req.Email)

// 	if !exists {
// 		return "", errors.New("DEBUG: user was not found")
// 	}

// 	println("DEBUG: User found:", user.Email)
// 	println("DEBUG: Password hash:", user.PasswordHash)

// 	err := utils.CheckPassword(user.PasswordHash, req.Password)

// 	if err != nil {
// 		println("DEBUG: bcrypt password check FAILED")
// 		return "", errors.New("invalid email or password")
// 	}

// 	println("DEBUG: bcrypt password check PASSED")

// 	return "Login successful", nil
// }
