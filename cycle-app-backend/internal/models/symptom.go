package models

type Symptom struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Date     string `json:"date"`
	Name     string `json:"name"`
	Severity int    `json:"severity"`
	Notes    string `json:"notes,omitempty"`
}
