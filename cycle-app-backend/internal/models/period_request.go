package models

type CreatePeriodRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Flow      string `json:"flow"`
	Notes     string `json:"notes"`
}
