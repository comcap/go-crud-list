package test

import (
	"context"
	"meeting-room/service/validator"

	"github.com/stretchr/testify/suite"

	"meeting-room/service/util"
	"meeting-room/service/util/mocks"
)

type PackageTestSuite struct {
	suite.Suite
	ctx         context.Context
	uuid        util.UUID
	validator   validator.Validator
	companyRepo *mocks.Repository
	staffRepo   *mocks.Repository
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	var err error
	suite.uuid, err = util.NewUUID()
	suite.NoError(err)
	suite.companyRepo = &mocks.Repository{}
	suite.staffRepo = &mocks.Repository{}
	suite.validator = validator.New(suite.companyRepo, suite.staffRepo)
}
