package internal_test

import (
	"testing"

	"github.com/ryanmoran/libgenders/internal"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testSet(t *testing.T, context spec.G, it spec.S) {
	var Expect = NewWithT(t).Expect

	context("Union", func() {
		it("returns a new set that is the union of the equal-length sets", func() {
			left := internal.Set([]int{1, 2, 3, 4})
			right := internal.Set([]int{3, 4, 5, 6})

			Expect(left.Union(right)).To(Equal(internal.Set([]int{1, 2, 3, 4, 5, 6})))
		})

		it("returns a new set that is the union of the unequal-length sets", func() {
			left := internal.Set([]int{1, 2, 3})
			right := internal.Set([]int{3, 4, 5, 6, 7})

			Expect(left.Union(right)).To(Equal(internal.Set([]int{1, 2, 3, 4, 5, 6, 7})))
		})

		it("returns a new set that is the union of the given sets", func() {
			left := internal.Set([]int{1, 2, 3, 4})
			right := internal.Set([]int{})

			Expect(left.Union(right)).To(Equal(internal.Set([]int{1, 2, 3, 4})))
		})
	})

	context("Intersection", func() {
		it("returns a new set that is the intersection of the equal-length sets", func() {
			left := internal.Set([]int{1, 2, 3, 4})
			right := internal.Set([]int{3, 4, 5, 6})

			Expect(left.Intersection(right)).To(Equal(internal.Set([]int{3, 4})))
		})

		it("returns a new set that is the intersection of the unequal-length sets", func() {
			left := internal.Set([]int{1, 2, 3})
			right := internal.Set([]int{3, 4, 5, 6, 7})

			Expect(left.Intersection(right)).To(Equal(internal.Set([]int{3})))
		})

		it("returns a new set that is the intersection of the given sets", func() {
			left := internal.Set([]int{1, 2, 3, 4})
			right := internal.Set([]int{})

			Expect(left.Intersection(right)).To(BeNil())
		})
	})

	context("Difference", func() {
		it("returns a new set that is the difference of the equal-length sets", func() {
			left := internal.Set([]int{1, 2, 3, 4})
			right := internal.Set([]int{3, 4, 5, 6})

			Expect(left.Difference(right)).To(Equal(internal.Set([]int{1, 2})))
		})

		it("returns a new set that is the difference of the unequal-length sets", func() {
			left := internal.Set([]int{1, 2, 3})
			right := internal.Set([]int{3, 4, 5, 6, 7})

			Expect(left.Difference(right)).To(Equal(internal.Set([]int{1, 2})))
		})

		it("returns a new set that is the difference of the given sets", func() {
			left := internal.Set([]int{1, 2, 3, 4})
			right := internal.Set([]int{})

			Expect(left.Difference(right)).To(Equal(internal.Set([]int{1, 2, 3, 4})))
		})
	})
}
