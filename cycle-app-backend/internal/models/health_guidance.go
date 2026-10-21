package models

type HealthGuidance struct {
	Symptom             string   `json:"symptom"`
	PossibleCauses      []string `json:"possible_causes"`
	SelfCare             []string `json:"self_care"`
	WhenToSeeGynecologist string   `json:"when_to_see_gynecologist"`
	UrgentWarning        string   `json:"urgent_warning"`
	Disclaimer           string   `json:"disclaimer"`
}