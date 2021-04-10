package test

import "meeting-room/service/company/companyin"

func (suite *PackageTestSuite) TestCreateInputToCompanyDomain() {
	given := companyin.MakeTestCreateInput()

	actual := companyin.CreateInputToCompanyDomain(given)

	suite.Equal(given.ID, actual.ID)
	suite.Equal(given.Name, actual.Name)
}
