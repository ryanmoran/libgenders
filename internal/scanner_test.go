package internal_test

import (
	"testing"

	"github.com/ryanmoran/libgenders/internal"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testScanner(t *testing.T, context spec.G, it spec.S) {
	var Expect = NewWithT(t).Expect

	context("Next", func() {
		it("splits a given query into tokenizable characters", func() {
			var parts []string
			scanner := internal.NewScanner("some query")
			for scanner.Len() > 0 {
				part, err := scanner.Next()
				Expect(err).NotTo(HaveOccurred())
				parts = append(parts, part)
			}

			Expect(parts).To(Equal([]string{"s", "o", "m", "e", " ", "q", "u", "e", "r", "y"}))
		})

		it("treats set operators specially", func() {
			var parts []string
			scanner := internal.NewScanner("w && x || y -- z")
			for scanner.Len() > 0 {
				part, err := scanner.Next()
				Expect(err).NotTo(HaveOccurred())
				parts = append(parts, part)
			}

			Expect(parts).To(Equal([]string{"w", " ", "&&", " ", "x", " ", "||", " ", "y", " ", "--", " ", "z"}))
		})
	})
}
