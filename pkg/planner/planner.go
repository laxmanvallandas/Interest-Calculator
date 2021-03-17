package planner

import (
	"errors"
	"strconv"
	"time"

	"github.com/laxmanvallandas/assignment/pkg/model"
)

//PlanRequest request payload
type PlanRequest struct {
	LoanAmount  string `json:"loanAmount"`
	NominalRate string `json:"nominalRate"`
	Duration    int64  `json:"duration"`
	StartDate   string `json:"startDate"`
}

// PlanResponse response payload
type PlanResponse struct {
	BorrowerPayments []model.BorrowerPayment `json:"borrowerPayments"`
}

const (
	ErrInvalidRequest = "received invalid request to generate plan"
)

func (p *PlanRequest) Plan() (PlanResponse, error) {
	var finalPlan []model.BorrowerPayment

	rate, err := strconv.ParseFloat(p.NominalRate, 64)
	if err != nil {
		return PlanResponse{}, err
	}

	plan, err := model.NewBorrowerPaymentPlan(p.LoanAmount, rate, p.Duration)
	if err != nil {
		return PlanResponse{}, err
	}

	startDate, _ := time.Parse(time.RFC3339, p.StartDate)

	// Loop duration(installments) number of times and form the borrowers montly plan
	for p.Duration > 0 {
		err := plan.GetMonthlyPlan(rate/100.0, startDate)
		if err != nil {
			return PlanResponse{}, err
		}
		startDate = nextAnnuityDate(startDate)
		finalPlan = append(finalPlan, *plan)
		p.Duration--
		plan.InitialOutstandingPrincipal = plan.RemainingOutstandingPrincipal
	}
	return PlanResponse{BorrowerPayments: finalPlan}, nil
}

func (p *PlanRequest) ValidateRequest() error {
	if p.LoanAmount == "" || p.NominalRate == "" || p.StartDate == "" || p.Duration <= 0 {
		return errors.New(ErrInvalidRequest)
	}
	return nil
}

func nextAnnuityDate(startDate time.Time) time.Time {
	nextStartDate := startDate.AddDate(0, 1, 0)

	year, nextMonth, _ := nextStartDate.Date()
	location := nextStartDate.Location()

	return time.Date(year, nextMonth, startDate.Day(), 0, 0, 0, 0, location)
}
