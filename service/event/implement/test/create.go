package test

import (
	"meeting-room/domain"
	"meeting-room/service/company/companyin"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestCreate() {
	givenInput := companyin.MakeTestCreateInput()
	givenCompany := domain.MakeTestCompany()

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.repo.On("Create", mock.Anything, givenCompany).Once().Return(givenCompany.ID, nil)
	actualID, err := suite.service.Create(suite.ctx, givenInput)

	suite.NoError(err)
	suite.Equal(givenCompany.ID, actualID)
}
