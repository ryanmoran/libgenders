package libgenders_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestLibgenders(t *testing.T) {
	suite := spec.New(" libgenders", spec.Report(report.Terminal{}))
	suite("Database", testDatabase)
	suite("Parser", testParser)
	suite("Node", testNode)
	suite.Run(t)
}
