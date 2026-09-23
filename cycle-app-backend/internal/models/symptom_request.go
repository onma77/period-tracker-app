package models

type CreateSymptomRequest struct {
	Date     string `json:"date"`
	Name     string `json:"name"`
	Severity int    `json:"severity"`
	Notes    string `json:"notes"`
}
