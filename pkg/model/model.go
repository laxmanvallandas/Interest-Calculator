package model

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// BorrowerPayment the payload in the generated plan
type BorrowerPayment struct {
	Amount                        string `json:"borrowerPaymentAmount"`
	Date                          string `json:"date"`
	InitialOutstandingPrincipal   string `json:"initialOutstandingPrincipal"`
	Interest                      string `json:"interest"`
	Principal                     string `json:"principal"`
	RemainingOutstandingPrincipal string `json:"remainingOutstandingPrincipal"`
}

const (
	daysInMonth = 30.0
	daysInYear  = 360.0
)

// NewBorrowerPaymentPlan returns BorrowerPayment with initial month values
func NewBorrowerPaymentPlan(loanAmount string, rate float64, noOfMonths int64) (*BorrowerPayment, error) {
	plan := &BorrowerPayment{
		InitialOutstandingPrincipal: loanAmount,
	}
	amount, err := strconv.ParseFloat(loanAmount, 64)
	if err != nil {
		return nil, err
	}
	plan.Amount = plan.getAnnuity(rate, amount, noOfMonths)

	return plan, nil
}

// GetMonthlyPlan gets the calculated montly plan
func (p *BorrowerPayment) GetMonthlyPlan(rate float64, startDate time.Time) error {
	p.Date = startDate.String()

	initialOutstandingPrincipal, err := strconv.ParseFloat(p.InitialOutstandingPrincipal, 64)
	if err != nil {
		return err
	}

	p.calculateInterest(rate, initialOutstandingPrincipal)

	if err := p.calculatePrincipal(rate); err != nil {
		return err
	}

	if err := p.calculateRemainingOutstandingPrincipal(initialOutstandingPrincipal); err != nil {
		return err
	}

	return nil
}

func (p *BorrowerPayment) getAnnuity(rate, initialOutstandingPrincipal float64, numberOfMonths int64) string {
	r := rate / (12 * 100.0)
	emi := (initialOutstandingPrincipal * r * math.Pow(1+r, float64(numberOfMonths))) / (math.Pow(1+r, float64(numberOfMonths)) - 1)

	return fmt.Sprintf("%.2f", emi)
}

func (p *BorrowerPayment) calculateInterest(rate, initialOutstandingPrincipal float64) {
	p.Interest = fmt.Sprintf("%.2f", (rate*daysInMonth*initialOutstandingPrincipal)/daysInYear)
}

func (p *BorrowerPayment) calculatePrincipal(rate float64) error {
	amount, err := strconv.ParseFloat(p.Amount, 64)
	if err != nil {
		return err
	}

	interest, err := strconv.ParseFloat(p.Interest, 64)
	if err != nil {
		return err
	}

	principal := amount - interest
	initialOutstandingPrincipal, err := strconv.ParseFloat(p.InitialOutstandingPrincipal, 64)
	if err != nil {
		return err
	}

	//  calculated principal amount exceeds the initial outstanding principal amount, take initial outstanding principal amount instead
	//  can happen in the very last installment
	if principal > initialOutstandingPrincipal {
		principal = initialOutstandingPrincipal
		// Last month Annuity would be recalculated based on current outstanding Principal
		p.Amount = p.getAnnuity(rate, principal, 1)
	}
	p.Principal = fmt.Sprintf("%.2f", principal)

	return nil
}

func (p *BorrowerPayment) calculateRemainingOutstandingPrincipal(initialOutstandingPrincipal float64) error {
	principal, err := strconv.ParseFloat(p.Principal, 64)
	if err != nil {
		return err
	}
	p.RemainingOutstandingPrincipal = fmt.Sprintf("%.2f", initialOutstandingPrincipal-principal)

	return nil
}
