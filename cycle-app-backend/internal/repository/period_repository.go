package repository

import "cycle-app-backend/internal/models"

type PeriodRepository struct {
	periods []models.Period
	nextID  int
}

func NewPeriodRepository() *PeriodRepository {
	return &PeriodRepository{
		periods: []models.Period{},
		nextID:  1,
	}
}

func (r *PeriodRepository) Create(period models.Period) models.Period {
	period.ID = r.nextID
	r.nextID++

	r.periods = append(r.periods, period)

	return period
}

func (r *PeriodRepository) GetByUserID(userID int) []models.Period {
	var userPeriods []models.Period

	for _, period := range r.periods {
		if period.UserID == userID {
			userPeriods = append(userPeriods, period)
		}
	}

	return userPeriods
}
