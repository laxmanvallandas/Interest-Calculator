package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/laxmanvallandas/assignment/pkg/planner"
)

// GeneratePlan handler to generate the plan
func GeneratePlan(w http.ResponseWriter, r *http.Request) {
	var planner planner.PlanRequest
	if err := json.NewDecoder(r.Body).Decode(&planner); err != nil {
		fmt.Println(err)
		return
	}

	err := planner.ValidateRequest()
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err)
		return
	}
	planResponse, err := planner.Plan()
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := writeResponse(w, http.StatusOK, planResponse); err != nil {
		fmt.Println("failed to write JSON response")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// WriteResponse encodes the json response and writes to http writer
func writeResponse(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
