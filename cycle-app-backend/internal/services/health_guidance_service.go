package services

import (
	"strings"

	"cycle-app-backend/internal/models"
)

type HealthGuidanceService struct{}

func NewHealthGuidanceService() *HealthGuidanceService {
	return &HealthGuidanceService{}
}

func (s *HealthGuidanceService) GetGuidance(symptomName string) models.HealthGuidance {
	symptom := strings.ToLower(strings.TrimSpace(symptomName))

	switch symptom {
	case "cramps", "menstrual cramps", "period pain":
		return models.HealthGuidance{
			Symptom: symptomName,
			PossibleCauses: []string{
				"Menstrual cramps",
				"Conditions such as endometriosis or adenomyosis",
				"Other pelvic conditions",
			},
			SelfCare: []string{
				"Apply a warm compress to your lower abdomen",
				"Rest and try gentle movement if comfortable",
				"Drink water and get enough sleep",
			},
			WhenToSeeGynecologist: "Arrange a gynecologist appointment if the pain is severe, worsening, disrupts daily activities, or does not improve with usual care.",
			UrgentWarning:         "Seek urgent medical care for sudden or severe pelvic pain, fainting, fever with pelvic pain, or heavy bleeding accompanied by dizziness.",
			Disclaimer:            "This guidance is educational and cannot identify the cause of your pain.",
		}

	case "heavy bleeding", "heavy period":
		return models.HealthGuidance{
			Symptom: symptomName,
			PossibleCauses: []string{
				"Menstrual or hormonal changes",
				"Conditions such as fibroids or adenomyosis",
				"Other reproductive health conditions",
			},
			SelfCare: []string{
				"Track how often you change pads or other period products",
				"Record how long the bleeding lasts",
				"Drink fluids and rest when needed",
			},
			WhenToSeeGynecologist: "Arrange an appointment if bleeding repeatedly lasts longer than seven days, affects daily life, or seems heavier than usual.",
			UrgentWarning:         "Seek urgent medical care if you soak through a pad or tampon every hour for two consecutive hours, especially with dizziness, weakness, shortness of breath, or fainting.",
			Disclaimer:            "Heavy bleeding can have different causes; this guidance is not a diagnosis.",
		}

	case "bloating":
		return models.HealthGuidance{
			Symptom: symptomName,
			PossibleCauses: []string{
				"Normal hormonal changes around menstruation",
				"Digestive issues",
				"Other conditions if bloating is persistent",
			},
			SelfCare: []string{
				"Drink enough water",
				"Try gentle movement",
				"Track when bloating occurs and whether anything makes it worse",
			},
			WhenToSeeGynecologist: "Seek medical advice if bloating is persistent, worsening, unusual for you, or occurs with ongoing pelvic pain or feeling full quickly.",
			UrgentWarning:         "Seek urgent care for severe abdominal pain, repeated vomiting, or a swollen abdomen with inability to pass stool or gas.",
			Disclaimer:            "Symptoms alone cannot determine the underlying cause.",
		}

	default:
		return models.HealthGuidance{
			Symptom: symptomName,
			PossibleCauses: []string{
				"Symptoms can have several possible causes",
			},
			SelfCare: []string{
				"Record when the symptom starts and how severe it feels",
				"Rest and drink enough water",
				"Avoid self-medicating if you are unsure which medicine is appropriate",
			},
			WhenToSeeGynecologist: "Speak with a healthcare professional if the symptom persists, worsens, recurs, or affects your daily activities.",
			UrgentWarning:         "Seek urgent medical care for sudden or severe pain, fainting, or heavy bleeding with dizziness.",
			Disclaimer:            "This guidance is educational and does not replace professional medical advice.",
		}
	}
}
