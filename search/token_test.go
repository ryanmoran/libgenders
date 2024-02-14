package search_test

import (
	"testing"

	"github.com/ryanmoran/libgenders/search"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testToken(t *testing.T, context spec.G, it spec.S) {
	var Expect = NewWithT(t).Expect

	context("Tokenize", func() {
		it("parses values", func() {
			tokens, err := search.Tokenize("this_is_a_value")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]search.Token{
				{Kind: search.ValueTokenKind, Text: "this_is_a_value"},
			}))
		})

		it("parses unions", func() {
			tokens, err := search.Tokenize("attr1 || attr2")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]search.Token{
				{Kind: search.UnionTokenKind, Text: "||"},
				{Kind: search.ValueTokenKind, Text: "attr2"},
				{Kind: search.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses intersections", func() {
			tokens, err := search.Tokenize("attr1 && attr2")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]search.Token{
				{Kind: search.IntersectionTokenKind, Text: "&&"},
				{Kind: search.ValueTokenKind, Text: "attr2"},
				{Kind: search.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses differences", func() {
			tokens, err := search.Tokenize("attr1 -- attr2")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]search.Token{
				{Kind: search.DifferenceTokenKind, Text: "--"},
				{Kind: search.ValueTokenKind, Text: "attr2"},
				{Kind: search.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses complements", func() {
			tokens, err := search.Tokenize("~attr1")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]search.Token{
				{Kind: search.ComplementTokenKind, Text: "~"},
				{Kind: search.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses parentheses", func() {
			tokens, err := search.Tokenize("(attr1)")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]search.Token{
				{Kind: search.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses mixed queries", func() {
			tokens, err := search.Tokenize("((attr1 && ~attr3) || (attr1 -- attr5)) && attr7")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]search.Token{
				{Kind: search.IntersectionTokenKind, Text: "&&"},
				{Kind: search.ValueTokenKind, Text: "attr7"},
				{Kind: search.UnionTokenKind, Text: "||"},
				{Kind: search.DifferenceTokenKind, Text: "--"},
				{Kind: search.ValueTokenKind, Text: "attr5"},
				{Kind: search.ValueTokenKind, Text: "attr1"},
				{Kind: search.IntersectionTokenKind, Text: "&&"},
				{Kind: search.ComplementTokenKind, Text: "~"},
				{Kind: search.ValueTokenKind, Text: "attr3"},
				{Kind: search.ValueTokenKind, Text: "attr1"},
			}))
		})
	})
}
