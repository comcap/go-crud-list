package test

import (
	"context"
	"meeting-room/service/util/mocks"
	"meeting-room/service/validator"

	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	ctx            context.Context
	repo           *mocks.Repository
	submissionRepo *mocks.Repository
	validator      validator.Validator
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.repo = &mocks.Repository{}
	suite.submissionRepo = &mocks.Repository{}
	suite.validator = validator.New(suite.repo, suite.submissionRepo)
}

type SimpleTestStruct struct {
	Title string `validate:"required"`
	Body  string `validate:"max=5"`
}

func (suite *PackageTestSuite) TestValidatorValidateValid() {
	given := &SimpleTestStruct{
		Title: "test",
		Body:  "AAA",
	}

	err := suite.validator.Validate(given)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestValidatorValidateInvalid() {
	given := &SimpleTestStruct{
		Title: "",
		Body:  "AAAAAAA",
	}

	err := suite.validator.Validate(given)

	suite.Error(err)
}
