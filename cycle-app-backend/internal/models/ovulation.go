package models

type OvulationPrediction struct {
	LastPeriodStart    string  `json:"last_period_start"`
	AverageCycleLength float64 `json:"average_cycle_length"`
	EstimatedOvulation string  `json:"estimated_ovulation"`
	FertileWindowStart string  `json:"fertile_window_start"`
	FertileWindowEnd   string  `json:"fertile_window_end"`
}
