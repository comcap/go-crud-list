package test

import "meeting-room/service/company/companyin"

func (suite *PackageTestSuite) TestUpdateInputToCompanyDomain() {
	given := companyin.MakeTestUpdateInput()

	actual := companyin.UpdateInputToCompanyDomain(given)

	suite.Equal(given.ID, actual.ID)
	suite.Equal(given.Name, actual.Name)
}
