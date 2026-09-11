package handlers

import (
	"encoding/json"
	"net/http"

	"cycle-app-backend/internal/middleware"
	"cycle-app-backend/internal/models"
	"cycle-app-backend/internal/services"
)

type CycleHandler struct {
	PeriodService *services.PeriodService
	CycleService  *services.CycleService
}

func NewCycleHandler(
	periodService *services.PeriodService,
	cycleService *services.CycleService,
) *CycleHandler {
	return &CycleHandler{
		PeriodService: periodService,
		CycleService:  cycleService,
	}
}

func (h *CycleHandler) GetCycle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	user, ok := r.Context().Value(
		middleware.UserContextKey(),
	).(map[string]interface{})

	if !ok {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	userID, ok := user["user_id"].(int)

	if !ok {
		http.Error(w, "Invalid user information", http.StatusUnauthorized)
		return
	}

	periods := h.PeriodService.GetUserPeriods(userID)

	if len(periods) == 0 {
		http.Error(
			w,
			"Not enough period data",
			http.StatusBadRequest,
		)
		return
	}

	averageCycle, err := h.CycleService.CalculateAverageCycle(periods)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lastPeriod := periods[0]

	for _, period := range periods {
		if period.StartDate > lastPeriod.StartDate {
			lastPeriod = period
		}
	}

	response := models.CycleSummary{
		AverageCycleLength: averageCycle,
		LastPeriodStart:    lastPeriod.StartDate,
		TotalPeriods:       len(periods),
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func (h *CycleHandler) GetPrediction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	user, ok := r.Context().Value(middleware.UserContextKey()).(map[string]interface{})
	if !ok {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	userID, ok := user["user_id"].(int)
	if !ok {
		http.Error(w, "Invalid user information", http.StatusUnauthorized)
		return
	}

	periods := h.PeriodService.GetUserPeriods(userID)

	if len(periods) < 2 {
		http.Error(w, "At least two periods are required", http.StatusBadRequest)
		return
	}

	averageCycle, err := h.CycleService.CalculateAverageCycle(periods)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lastPeriod := periods[0]

	for _, period := range periods {
		if period.StartDate > lastPeriod.StartDate {
			lastPeriod = period
		}
	}

	estimatedNextPeriod, err := h.CycleService.PredictNextPeriod(periods)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := models.CyclePrediction{
		LastPeriodStart:     lastPeriod.StartDate,
		AverageCycleLength:  averageCycle,
		EstimatedNextPeriod: estimatedNextPeriod,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
