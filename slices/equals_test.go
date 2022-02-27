package slices_test

import (
	"github.com/boundedinfinity/commons/slices"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("EqualsFn()", func() {
	Context("of string slice", func() {
		fn := func(x, y string) bool { return x == y }

		It("should be equal", func() {
			xs := []string{"a", "b", "c"}
			ys := []string{"a", "b", "c"}
			Expect(slices.EqualsFn(xs, ys, fn)).To(BeTrue())
		})

		It("should not be equal", func() {
			xs := []string{"a", "b", "c"}
			ys := []string{"b", "c", "e"}
			Expect(slices.EqualsFn(xs, ys, fn)).To(BeFalse())
		})
	})

	Context("of int slice", func() {
		fn := func(x, y int) bool { return x == y }

		It("should be equal", func() {
			xs := []int{1, 2, 3}
			ys := []int{1, 2, 3}
			Expect(slices.EqualsFn(xs, ys, fn)).To(BeTrue())
		})

		It("should not be equal", func() {
			xs := []int{1, 2, 3}
			ys := []int{2, 3, 4}
			Expect(slices.EqualsFn(xs, ys, fn)).To(BeFalse())
		})
	})
})

var _ = Describe("Equals()", func() {
	Context("of string slice", func() {
		It("should be equal", func() {
			xs := []string{"a", "b", "c"}
			ys := []string{"a", "b", "c"}
			Expect(slices.Equals(xs, ys)).To(BeTrue())
		})

		It("should not be equal", func() {
			xs := []string{"a", "b", "c"}
			ys := []string{"b", "c", "e"}
			Expect(slices.Equals(xs, ys)).To(BeFalse())
		})
	})

	Context("of int slice", func() {
		It("should be equal", func() {
			xs := []int{1, 2, 3}
			ys := []int{1, 2, 3}
			Expect(slices.Equals(xs, ys)).To(BeTrue())
		})

		It("should not be equal", func() {
			xs := []int{1, 2, 3}
			ys := []int{2, 3, 4}
			Expect(slices.Equals(xs, ys)).To(BeFalse())
		})
	})
})
