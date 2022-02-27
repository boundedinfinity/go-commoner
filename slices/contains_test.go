package slices_test

import (
	"github.com/boundedinfinity/commons/slices"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Any()", func() {
	Context("of string slice", func() {
		It("should be found", func() {
			xs := []string{"a", "b", "c"}
			Expect(slices.Contains(xs, "a")).To(BeTrue())
		})

		It("should be found", func() {
			xs := []string{"a", "b", "c"}
			Expect(slices.Contains(xs, "A")).To(BeFalse())
		})
	})

	Context("of int slice", func() {
		It("should be equal", func() {
			xs := []int{1, 2, 3}
			Expect(slices.Contains(xs, 1)).To(BeTrue())
		})

		It("should not be equal", func() {
			xs := []int{1, 2, 3}
			Expect(slices.Contains(xs, 4)).To(BeFalse())
		})
	})
})
