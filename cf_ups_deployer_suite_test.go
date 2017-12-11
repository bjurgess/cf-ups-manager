package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCfUpsDeployer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CfUpsDeployer Suite")
}
