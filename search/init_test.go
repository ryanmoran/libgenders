package search_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestSearch(t *testing.T) {
	suite := spec.New(" libgenders/search", spec.Report(report.Terminal{}))
	suite("Query", testQuery)
	suite.Pend("Scanner", testScanner)
	suite.Pend("Stack", testStack)
	suite("Set", testSet)
	suite("Token", testToken)
	suite.Run(t)
}
