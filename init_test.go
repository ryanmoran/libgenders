package libgenders_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestLibgenders(t *testing.T) {
	suite := spec.New(" libgenders", spec.Report(report.Terminal{}))
	suite("Database", testDatabase)
	suite("Node", testNode)
	suite("Parser", testParser)
	suite("QueryEngine", testQueryEngine)
	suite.Run(t)
}
