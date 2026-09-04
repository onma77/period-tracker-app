package models

//import "time"

type User struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	DateOfBirth  string `json:"date_of_birth"`

	// ID              int64     `json:"id"`
	// FirstName       string    `json:"first_name"`
	// LastName        string    `json:"last_name"`
	// Email           string    `json:"email"`
	// Password        string    `json:"password"`
	// DateOfBirth     time.Time `json:"date_of_birth"`
	// CycleLength     int       `json:"cycle_length"`
	// PeriodLength    int       `json:"period_length"`
	// LastPeriodStart time.Time `json:"last_period_start"`
	// CreatedAt       time.Time `json:"created_at"`
}
