package models

type Period struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date,omitempty"`
	Flow      string `json:"flow,omitempty"`
	Notes     string `json:"notes,omitempty"`
}
