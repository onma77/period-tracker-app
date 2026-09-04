package models

type CycleSummary struct {
	AverageCycleLength float64 `json:"average_cycle_length"`
	LastPeriodStart    string  `json:"last_period_start"`
	TotalPeriods       int     `json:"total_periods"`
}
