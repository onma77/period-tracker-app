package services

import (
	"errors"

	"cycle-app-backend/internal/models"
	"cycle-app-backend/internal/utils"
)

type CycleService struct{}

func NewCycleService() *CycleService {
	return &CycleService{}
}

func (s *CycleService) CalculateCycleLength(
	current models.Period,
	previous models.Period,
) (int, error) {

	currentStart, err := utils.ParseDate(current.StartDate)
	if err != nil {
		return 0, err
	}

	previousStart, err := utils.ParseDate(previous.StartDate)
	if err != nil {
		return 0, err
	}

	if !currentStart.After(previousStart) {
		return 0, errors.New("current period must be after previous period")
	}

	return int(currentStart.Sub(previousStart).Hours() / 24), nil
}

func (s *CycleService) CalculatePeriodDuration(
	period models.Period,
) (int, error) {

	if period.EndDate == "" {
		return 0, errors.New("end date is required to calculate period duration")
	}

	startDate, err := utils.ParseDate(period.StartDate)
	if err != nil {
		return 0, err
	}

	endDate, err := utils.ParseDate(period.EndDate)
	if err != nil {
		return 0, err
	}

	if endDate.Before(startDate) {
		return 0, errors.New("end date cannot be before start date")
	}

	// +1 because the first day counts as a period day.
	return int(endDate.Sub(startDate).Hours()/24) + 1, nil
}
