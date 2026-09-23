package repository

import "cycle-app-backend/internal/models"

type SymptomRepository struct {
	symptoms []models.Symptom
	nextID   int
}

func NewSymptomRepository() *SymptomRepository {
	return &SymptomRepository{
		symptoms: []models.Symptom{},
		nextID:   1,
	}
}

func (r *SymptomRepository) Create(symptom models.Symptom) models.Symptom {
	symptom.ID = r.nextID
	r.nextID++

	r.symptoms = append(r.symptoms, symptom)

	return symptom
}

func (r *SymptomRepository) GetByUserID(userID int) []models.Symptom {
	var userSymptoms []models.Symptom

	for _, symptom := range r.symptoms {
		if symptom.UserID == userID {
			userSymptoms = append(userSymptoms, symptom)
		}
	}

	return userSymptoms
}
