package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/laxmanvallandas/assignment/pkg/planner"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePlan(t *testing.T) {
	data := `{"loanAmount": "5000","nominalRate": "5.0","duration": 24,"startDate": "2020-11-01T00:00:01Z"}`
	req, err := http.NewRequest("POST", "/generate-plan", strings.NewReader(data))
	assert.Nil(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	GeneratePlan(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)

	var pResp planner.PlanResponse
	err = json.NewDecoder(resp.Body).Decode(&pResp)
	assert.Nil(t, err)

	t.Logf("resp %+v", pResp)
	tmpData, _ := json.MarshalIndent(pResp, "", "    ")
	t.Logf("resp %s", string(tmpData))
}
