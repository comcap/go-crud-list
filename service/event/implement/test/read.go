package test

import (
	"meeting-room/domain"
	"meeting-room/service/company/companyin"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestRead() {
	givenInput := companyin.MakeTestReadInput()
	givenCompanyIDFilter := suite.makeCompanyIDFilter(givenInput.CompanyID)
	givenCompany := &domain.Company{}

	suite.repo.On("Read", mock.Anything, givenCompanyIDFilter, givenCompany).Once().Return(nil)
	actualView, err := suite.service.Read(suite.ctx, givenInput)

	suite.NoError(err)
	suite.Equal(givenCompany.ID, actualView.ID)
	suite.Equal(givenCompany.Name, actualView.Name)
}
