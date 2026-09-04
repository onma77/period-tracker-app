package services

import (
	"errors"
	"sort"

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

	return utils.DaysBetween(previousStart, currentStart), nil
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

	return utils.DaysBetween(startDate, endDate) + 1, nil
}

func (s *CycleService) CalculateAverageCycle(
	periods []models.Period,
) (float64, error) {

	if len(periods) < 2 {
		return 0, errors.New("at least two periods are required")
	}

	sort.Slice(periods, func(i, j int) bool {
		return periods[i].StartDate < periods[j].StartDate
	})

	totalDays := 0
	cycleCount := 0

	for i := 1; i < len(periods); i++ {

		cycleLength, err := s.CalculateCycleLength(
			periods[i],
			periods[i-1],
		)

		if err != nil {
			return 0, err
		}

		totalDays += cycleLength
		cycleCount++
	}

	return float64(totalDays) / float64(cycleCount), nil
}
