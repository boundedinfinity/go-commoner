package strings_test

import (
	"github.com/boundedinfinity/commons/strings"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Strings", func() {
	Context("Trim()", func() {
		It("trim normal string with one space", func() {
			input := " a string "
			expected := "a string"
			actual := strings.TrimSpace(input)

			Expect(actual).To(Equal(expected))
		})

		It("trim normal string with multiple spaces", func() {
			input := "   a string   "
			expected := "a string"
			actual := strings.TrimSpace(input)

			Expect(actual).To(Equal(expected))
		})
	})

	Context("Trim()", func() {
		It("trim normal string of ! and ^", func() {
			input := "!^a string^!"
			expected := "a string"
			actual := strings.Trim(input, "!^")

			Expect(actual).To(Equal(expected))
		})
	})

	Context("TrimFn()", func() {
		It("trim normal string with one space", func() {
			input := " a string "
			expected := "a string"
			actual := strings.TrimFunc(input, func(x rune) bool { return x == ' ' })

			Expect(actual).To(Equal(expected))
		})

		It("trim normal string with multiple spaces", func() {
			input := "   a string   "
			expected := "a string"
			actual := strings.TrimFunc(input, func(x rune) bool { return x == ' ' })

			Expect(actual).To(Equal(expected))
		})
	})
})
