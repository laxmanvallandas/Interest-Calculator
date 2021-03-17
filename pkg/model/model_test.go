package model_test

import (
	"testing"
	"time"

	"github.com/laxmanvallandas/assignment/pkg/model"
	"github.com/stretchr/testify/suite"
)

type modelSuite struct {
	suite.Suite
}

const (
	rate = 0.05
	//start                         = "2020-11-05T00:00:01Z"
	borrowerAmount                = "219.36"
	initialOutstandingPrincipal   = "5000"
	interest                      = "20.83"
	principal                     = "198.53"
	remainingOutstandingPrincipal = "4801.47"
)

func TestRunModelSuite(t *testing.T) {
	suite.Run(t, new(modelSuite))
}

func (m *modelSuite) TestGetplan() {
	plan, err := model.NewBorrowerPaymentPlan("5000", 5.0, 24)
	m.NoError(err)

	startDate, _ := time.Parse(time.RFC3339, "2020-11-05T00:00:01Z")
	err = plan.GetMonthlyPlan(rate, startDate)
	m.NoError(err)
	m.Equal(borrowerAmount, plan.Amount)
	m.Equal(initialOutstandingPrincipal, plan.InitialOutstandingPrincipal)
	m.Equal(interest, plan.Interest)
	m.Equal(principal, plan.Principal)
	m.Equal(remainingOutstandingPrincipal, plan.RemainingOutstandingPrincipal)
}

func (m *modelSuite) TestGetplanFailWhileModelInit() {
	_, err := model.NewBorrowerPaymentPlan("a", 5.0, 24)
	m.Error(err)
}

func (m *modelSuite) TestGetplanFailDueToInvalidInitialPrincipal() {
	plan, err := model.NewBorrowerPaymentPlan("5000", 5.0, 24)
	m.NoError(err)

	startDate, _ := time.Parse(time.RFC3339, "2020-11-05T00:00:01Z")

	plan.InitialOutstandingPrincipal = "a"
	err = plan.GetMonthlyPlan(rate, startDate)
	m.Error(err)
}
