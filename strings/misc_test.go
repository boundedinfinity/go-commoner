package strings_test

import (
	"github.com/boundedinfinity/commons/strings"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Strings", func() {
	Context("IsBlank()", func() {
		It("should find empty from normal string", func() {
			Expect(strings.IsBlank("")).To(BeTrue())
		})

		It("should find empty from typed string", func() {
			type MyString string
			var ms MyString = ""
			Expect(strings.IsBlank(ms)).To(BeTrue())
		})
	})

	Context("IsEmpty()", func() {
		It("should find empty from normal string", func() {
			Expect(strings.IsEmpty("")).To(BeTrue())
		})

		It("should find empty from typed string", func() {
			type MyString string
			var ms MyString = ""
			Expect(strings.IsEmpty(ms)).To(BeTrue())
		})
	})

	Context("Abbreviate()", func() {
		It("a string that is too long", func() {
			input := "a test string that is too long"
			expected := "a test string..."
			actual := strings.Abbreviate(input, 16)

			Expect(actual).To(Equal(expected))
		})

		It("not a string that is shorter", func() {
			input := "a test string"
			expected := "a test string"
			actual := strings.Abbreviate(input, 16)

			Expect(actual).To(Equal(expected))
		})

		It("not a string that is equal to the width", func() {
			input := "a test string"
			expected := "a test ..."
			actual := strings.Abbreviate(input, len(input)-3)

			Expect(actual).To(Equal(expected))
		})
	})
})
