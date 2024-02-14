package internal_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestInternal(t *testing.T) {
	suite := spec.New(" libgenders/internal", spec.Report(report.Terminal{}))
	suite("Parser", testParser)
	suite("Query", testQuery)
	suite.Pend("Scanner", testScanner)
	suite.Pend("Stack", testStack)
	suite("Set", testSet)
	suite("Token", testToken)
	suite.Run(t)
}
