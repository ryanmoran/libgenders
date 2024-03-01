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
	suite("Scanner", testScanner)
	suite("Set", testSet)
	suite("Stack", testStack)
	suite("Token", testToken)
	suite.Run(t)
}
