package strings_test

import (
	"github.com/boundedinfinity/commons/strings"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Strings", func() {
	Context("Split()", func() {
		It("should find empty from normal string", func() {
			input := "a b c d"
			expected := []string{"a", "b", "c", "d"}
			actual := strings.Split(input)

			Expect(actual).To(Equal(expected))
		})

		It("should find empty from typed string", func() {
			type MyString string
			var input MyString = "a b c d"
			expected := []MyString{"a", "b", "c", "d"}
			actual := strings.Split(input)

			Expect(actual).To(Equal(expected))
		})
	})
})
