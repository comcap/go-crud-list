package implement_test

import (
	"testing"

	"meeting-room/service/company/implement/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
