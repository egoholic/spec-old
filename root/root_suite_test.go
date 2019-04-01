package root_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRoot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Root Suite")
}
