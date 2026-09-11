package models

type CyclePrediction struct {
	LastPeriodStart       string `json:"last_period_start"`
	AverageCycleLength    float64 `json:"average_cycle_length"`
	EstimatedNextPeriod   string `json:"estimated_next_period"`
}