package internal_test

import (
	"testing"

	"github.com/ryanmoran/libgenders/internal"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testToken(t *testing.T, context spec.G, it spec.S) {
	var Expect = NewWithT(t).Expect

	context("Tokenize", func() {
		it("parses values", func() {
			tokens, err := internal.Tokenize("this_is_a_value")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]internal.Token{
				{Kind: internal.ValueTokenKind, Text: "this_is_a_value"},
			}))
		})

		it("parses unions", func() {
			tokens, err := internal.Tokenize("attr1 || attr2")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]internal.Token{
				{Kind: internal.UnionTokenKind, Text: "||"},
				{Kind: internal.ValueTokenKind, Text: "attr2"},
				{Kind: internal.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses intersections", func() {
			tokens, err := internal.Tokenize("attr1 && attr2")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]internal.Token{
				{Kind: internal.IntersectionTokenKind, Text: "&&"},
				{Kind: internal.ValueTokenKind, Text: "attr2"},
				{Kind: internal.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses differences", func() {
			tokens, err := internal.Tokenize("attr1 -- attr2")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]internal.Token{
				{Kind: internal.DifferenceTokenKind, Text: "--"},
				{Kind: internal.ValueTokenKind, Text: "attr2"},
				{Kind: internal.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses complements", func() {
			tokens, err := internal.Tokenize("~attr1")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]internal.Token{
				{Kind: internal.ComplementTokenKind, Text: "~"},
				{Kind: internal.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses parentheses", func() {
			tokens, err := internal.Tokenize("(attr1)")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]internal.Token{
				{Kind: internal.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses mixed queries", func() {
			tokens, err := internal.Tokenize("((attr1 && ~attr3) || (attr1 -- attr5)) && attr7")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]internal.Token{
				{Kind: internal.IntersectionTokenKind, Text: "&&"},
				{Kind: internal.ValueTokenKind, Text: "attr7"},
				{Kind: internal.UnionTokenKind, Text: "||"},
				{Kind: internal.DifferenceTokenKind, Text: "--"},
				{Kind: internal.ValueTokenKind, Text: "attr5"},
				{Kind: internal.ValueTokenKind, Text: "attr1"},
				{Kind: internal.IntersectionTokenKind, Text: "&&"},
				{Kind: internal.ComplementTokenKind, Text: "~"},
				{Kind: internal.ValueTokenKind, Text: "attr3"},
				{Kind: internal.ValueTokenKind, Text: "attr1"},
			}))
		})
	})
}
