package slices_test

import (
	"github.com/boundedinfinity/commons/slices"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("All()", func() {
	Context("of string slice", func() {
		It("should be found", func() {
			expected := []string{"a", "a"}
			input := []string{"a", "b", "a"}
			actual := slices.All(input, "a")

			Expect(actual).To(Equal(expected))
		})
	})
})
