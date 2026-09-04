package services

import (
	"errors"

	"cycle-app-backend/internal/models"
	"cycle-app-backend/internal/repository"
	"cycle-app-backend/internal/utils"
)

type PeriodService struct {
	periodRepo *repository.PeriodRepository
}

func NewPeriodService(repo *repository.PeriodRepository) *PeriodService {
	return &PeriodService{
		periodRepo: repo,
	}
}

func (s *PeriodService) CreatePeriod(
	userID int,
	req models.CreatePeriodRequest,
) (models.Period, error) {

	if userID <= 0 {
		return models.Period{}, errors.New("invalid user")
	}

	startDate, err := utils.ParseDate(req.StartDate)
	if err != nil {
		return models.Period{}, err
	}

	var endDate string

	if req.EndDate != "" {
		parsedEndDate, err := utils.ParseDate(req.EndDate)
		if err != nil {
			return models.Period{}, err
		}

		if parsedEndDate.Before(startDate) {
			return models.Period{}, errors.New("end date cannot be before start date")
		}

		endDate = req.EndDate
	}

	period := models.Period{
		UserID:    userID,
		StartDate: req.StartDate,
		EndDate:   endDate,
		Flow:      req.Flow,
		Notes:     req.Notes,
	}

	return s.periodRepo.Create(period), nil
}

func (s *PeriodService) GetUserPeriods(userID int) []models.Period {
	return s.periodRepo.GetByUserID(userID)
}
