package planner_test

import (
	"testing"

	"github.com/laxmanvallandas/assignment/pkg/planner"
	"github.com/stretchr/testify/suite"
)

type plannerSuite struct {
	suite.Suite
}

func TestRunPlannerSuite(t *testing.T) {
	suite.Run(t, new(plannerSuite))
}

func (p *plannerSuite) TestValidateRequestSuccess() {
	planRequest := &planner.PlanRequest{
		LoanAmount:  "5000",
		NominalRate: "5.0",
		Duration:    12,
		StartDate:   "2020-11-05T00:00:01Z",
	}
	err := planRequest.ValidateRequest()
	p.NoError(err)
}

func (p *plannerSuite) TestValidateRequestFailed() {
	planRequest := &planner.PlanRequest{
		LoanAmount: "5000",
		Duration:   12,
		StartDate:  "2020-11-05T00:00:01Z",
	}
	err := planRequest.ValidateRequest()
	p.Error(err)
	p.EqualError(err, planner.ErrInvalidRequest)
}

func (p *plannerSuite) TestPlanSuccess() {
	planRequest := &planner.PlanRequest{
		LoanAmount:  "5000",
		NominalRate: "5.0",
		Duration:    12,
		StartDate:   "2020-11-05T00:00:01Z",
	}
	response, err := planRequest.Plan()
	p.NoError(err)
	p.Equal(len(response.BorrowerPayments), 12)
}

func (p *plannerSuite) TestPlanFailure() {
	planRequest := &planner.PlanRequest{
		LoanAmount:  "5000",
		NominalRate: "a",
		Duration:    12,
		StartDate:   "2020-11-05T00:00:01Z",
	}
	_, err := planRequest.Plan()
	p.Error(err)
	p.EqualError(err, "strconv.ParseFloat: parsing \"a\": invalid syntax")
}
