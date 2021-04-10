package test

import (
	"context"
	"fmt"

	"meeting-room/service/company"
	"meeting-room/service/company/implement"
	"meeting-room/service/util/mocks"
	validatorMocks "meeting-room/service/validator/mocks"

	"github.com/stretchr/testify/suite"
)

// PackageTestSuite struct
type PackageTestSuite struct {
	suite.Suite
	ctx       context.Context
	validator *validatorMocks.Validator
	repo      *mocks.Repository
	uuid      *mocks.UUID
	service   company.Service
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.validator = &validatorMocks.Validator{}
	suite.repo = &mocks.Repository{}
	suite.service = implement.New(suite.validator, suite.repo, suite.uuid)
}

func (suite *PackageTestSuite) makeCompanyIDFilter(companyID string) (filters []string) {
	return []string{
		fmt.Sprintf("id:eq:%s", companyID),
	}
}
