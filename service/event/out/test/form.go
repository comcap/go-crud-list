package test

import (
	"meeting-room/domain"
	"meeting-room/service/company/out"
)

func (suite *PackageTestSuite) TestCompanyToView() {
	given := domain.MakeTestCompany()

	actual := out.CompanyToView(given)

	suite.Equal(given.ID, actual.ID)
	suite.Equal(given.Name, actual.Name)
}
