package services

import (
	"errors"

	"cycle-app-backend/internal/models"
	"cycle-app-backend/internal/repository"
	"cycle-app-backend/internal/utils"
)

type SymptomService struct {
	symptomRepo *repository.SymptomRepository
}

func NewSymptomService(repo *repository.SymptomRepository) *SymptomService {
	return &SymptomService{
		symptomRepo: repo,
	}
}

func (s *SymptomService) CreateSymptom(
	userID int,
	req models.CreateSymptomRequest,
) (models.Symptom, error) {

	if userID <= 0 {
		return models.Symptom{}, errors.New("invalid user")
	}

	if req.Name == "" {
		return models.Symptom{}, errors.New("symptom name is required")
	}

	if req.Severity < 1 || req.Severity > 10 {
		return models.Symptom{}, errors.New("severity must be between 1 and 10")
	}

	_, err := utils.ParseDate(req.Date)
	if err != nil {
		return models.Symptom{}, err
	}

	symptom := models.Symptom{
		UserID:   userID,
		Date:     req.Date,
		Name:     req.Name,
		Severity: req.Severity,
		Notes:    req.Notes,
	}

	return s.symptomRepo.Create(symptom), nil
}

func (s *SymptomService) GetUserSymptoms(userID int) []models.Symptom {
	return s.symptomRepo.GetByUserID(userID)
}
