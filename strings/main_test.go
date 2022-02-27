package strings_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJsonSchema(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Strings Suite")
}
