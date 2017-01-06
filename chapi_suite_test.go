package chapi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCHAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CHAPI Suite")
}
