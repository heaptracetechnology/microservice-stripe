package BalanceOperation

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestBalance(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("../test-report/cireport.txt")
	RunSpecsWithDefaultAndCustomReporters(t, "Balance Operations", []Reporter{junitReporter})
}
