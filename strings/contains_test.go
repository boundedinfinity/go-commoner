package strings_test

import (
	"github.com/boundedinfinity/commons/strings"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Strings", func() {
	Context("Uppercase()", func() {
		It("a normal string", func() {
			input := "string"
			expected := "STRING"
			actual := strings.Uppercase(input)

			Expect(actual).To(Equal(expected))
		})
	})
})
