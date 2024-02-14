package libgenders_test

import (
	"testing"

	"github.com/ryanmoran/libgenders"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testToken(t *testing.T, context spec.G, it spec.S) {
	var Expect = NewWithT(t).Expect

	context("Tokenize", func() {
		it("parses values", func() {
			tokens, err := libgenders.Tokenize("this_is_a_value")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]libgenders.Token{
				{Kind: libgenders.ValueTokenKind, Text: "this_is_a_value"},
			}))
		})

		it("parses unions", func() {
			tokens, err := libgenders.Tokenize("attr1 || attr2")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]libgenders.Token{
				{Kind: libgenders.UnionTokenKind, Text: "||"},
				{Kind: libgenders.ValueTokenKind, Text: "attr2"},
				{Kind: libgenders.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses intersections", func() {
			tokens, err := libgenders.Tokenize("attr1 && attr2")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]libgenders.Token{
				{Kind: libgenders.IntersectionTokenKind, Text: "&&"},
				{Kind: libgenders.ValueTokenKind, Text: "attr2"},
				{Kind: libgenders.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses differences", func() {
			tokens, err := libgenders.Tokenize("attr1 -- attr2")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]libgenders.Token{
				{Kind: libgenders.DifferenceTokenKind, Text: "--"},
				{Kind: libgenders.ValueTokenKind, Text: "attr2"},
				{Kind: libgenders.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses complements", func() {
			tokens, err := libgenders.Tokenize("~attr1")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]libgenders.Token{
				{Kind: libgenders.ComplementTokenKind, Text: "~"},
				{Kind: libgenders.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses parentheses", func() {
			tokens, err := libgenders.Tokenize("(attr1)")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]libgenders.Token{
				{Kind: libgenders.ValueTokenKind, Text: "attr1"},
			}))
		})

		it("parses mixed queries", func() {
			tokens, err := libgenders.Tokenize("((attr1 && ~attr3) || (attr1 -- attr5)) && attr7")
			Expect(err).NotTo(HaveOccurred())
			Expect(tokens).To(Equal([]libgenders.Token{
				{Kind: libgenders.IntersectionTokenKind, Text: "&&"},
				{Kind: libgenders.ValueTokenKind, Text: "attr7"},
				{Kind: libgenders.UnionTokenKind, Text: "||"},
				{Kind: libgenders.DifferenceTokenKind, Text: "--"},
				{Kind: libgenders.ValueTokenKind, Text: "attr5"},
				{Kind: libgenders.ValueTokenKind, Text: "attr1"},
				{Kind: libgenders.IntersectionTokenKind, Text: "&&"},
				{Kind: libgenders.ComplementTokenKind, Text: "~"},
				{Kind: libgenders.ValueTokenKind, Text: "attr3"},
				{Kind: libgenders.ValueTokenKind, Text: "attr1"},
			}))
		})
	})
}
