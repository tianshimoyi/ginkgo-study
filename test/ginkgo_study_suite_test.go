package test_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgoStudy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoStudy Suite")
}
