package libgenders_test

import (
	"testing"

	"github.com/ryanmoran/libgenders"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testQuery(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		attrs = map[string]libgenders.Set{
			"attr1":  {0, 1, 2, 3, 4, 5, 6, 7},
			"attr2":  {0, 1, 2, 3, 4, 5, 6, 7},
			"attr3":  {0, 1, 2, 3},
			"attr4":  {0, 1, 2, 3},
			"attr5":  {4, 5, 6, 7},
			"attr6":  {4, 5, 6, 7},
			"attr7":  {0, 2, 4, 6},
			"attr8":  {0, 2, 4, 6},
			"attr9":  {1, 3, 5, 7},
			"attr10": {1, 3, 5, 7},
		}
		attrvals = map[string]libgenders.Set{
			"attr10=val10": {1, 3, 5, 7},
			"attr2=val2":   {0, 1, 2, 3, 4, 5, 6, 7},
			"attr4=val4":   {0, 1, 2, 3},
			"attr6=val6":   {4, 5, 6, 7},
			"attr8=val8":   {0, 2, 4, 6},
		}
		indices = libgenders.Set{0, 1, 2, 3, 4, 5, 6, 7}
	)

	context("ParseQuery", func() {
		it("parses attribute-only queries", func() {
			tokens := []libgenders.Token{
				{Kind: libgenders.ValueTokenKind, Text: "attr1"},
			}

			Expect(libgenders.ParseQuery(tokens)).To(Equal(libgenders.ValueQuery{
				Expression: "attr1",
			}))
		})

		it("parses attribute-value queries", func() {
			tokens := []libgenders.Token{
				{Kind: libgenders.ValueTokenKind, Text: "attr2=val2"},
			}

			Expect(libgenders.ParseQuery(tokens)).To(Equal(libgenders.ValueQuery{
				Expression: "attr2=val2",
			}))
		})

		it("parses union queries", func() {
			tokens := []libgenders.Token{
				{Kind: libgenders.UnionTokenKind, Text: "||"},
				{Kind: libgenders.ValueTokenKind, Text: "attr2"},
				{Kind: libgenders.ValueTokenKind, Text: "attr1"},
			}

			Expect(libgenders.ParseQuery(tokens)).To(Equal(libgenders.UnionQuery{
				Left:  libgenders.ValueQuery{Expression: "attr1"},
				Right: libgenders.ValueQuery{Expression: "attr2"},
			}))
		})

		it("parses intersection queries", func() {
			tokens := []libgenders.Token{
				{Kind: libgenders.IntersectionTokenKind, Text: "&&"},
				{Kind: libgenders.ValueTokenKind, Text: "attr2"},
				{Kind: libgenders.ValueTokenKind, Text: "attr1"},
			}

			Expect(libgenders.ParseQuery(tokens)).To(Equal(libgenders.IntersectionQuery{
				Left:  libgenders.ValueQuery{Expression: "attr1"},
				Right: libgenders.ValueQuery{Expression: "attr2"},
			}))
		})

		it("parses difference queries", func() {
			tokens := []libgenders.Token{
				{Kind: libgenders.DifferenceTokenKind, Text: "--"},
				{Kind: libgenders.ValueTokenKind, Text: "attr2"},
				{Kind: libgenders.ValueTokenKind, Text: "attr1"},
			}

			Expect(libgenders.ParseQuery(tokens)).To(Equal(libgenders.DifferenceQuery{
				Left:  libgenders.ValueQuery{Expression: "attr1"},
				Right: libgenders.ValueQuery{Expression: "attr2"},
			}))
		})

		it("parses complement queries", func() {
			tokens := []libgenders.Token{
				{Kind: libgenders.ComplementTokenKind, Text: "~"},
				{Kind: libgenders.ValueTokenKind, Text: "attr1"},
			}

			Expect(libgenders.ParseQuery(tokens)).To(Equal(libgenders.ComplementQuery{
				Query: libgenders.ValueQuery{Expression: "attr1"},
			}))
		})

		it("parses mixed queries", func() {
			tokens := []libgenders.Token{
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
			}

			Expect(libgenders.ParseQuery(tokens)).To(Equal(libgenders.IntersectionQuery{
				Left: libgenders.UnionQuery{
					Left: libgenders.IntersectionQuery{
						Left: libgenders.ValueQuery{Expression: "attr1"},
						Right: libgenders.ComplementQuery{
							Query: libgenders.ValueQuery{Expression: "attr3"},
						},
					},
					Right: libgenders.DifferenceQuery{
						Left:  libgenders.ValueQuery{Expression: "attr1"},
						Right: libgenders.ValueQuery{Expression: "attr5"},
					},
				},
				Right: libgenders.ValueQuery{Expression: "attr7"},
			}))
		})
	})

	context("ValueQuery", func() {
		it("returns a set matching the attribute-only query", func() {
			query := libgenders.ValueQuery{Expression: "attr1"}
			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(libgenders.Set{0, 1, 2, 3, 4, 5, 6, 7}))
		})

		it("returns a set matching the attribute-value query", func() {
			query := libgenders.ValueQuery{Expression: "attr4=val4"}
			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(libgenders.Set{0, 1, 2, 3}))
		})
	})

	context("UnionQuery", func() {
		it("returns a set matching the query", func() {
			query := libgenders.UnionQuery{
				Left:  libgenders.ValueQuery{Expression: "attr4"},
				Right: libgenders.ValueQuery{Expression: "attr8=val8"},
			}

			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(libgenders.Set{0, 1, 2, 3, 4, 6}))
		})
	})

	context("IntersectionQuery", func() {
		it("returns a set matching the query", func() {
			query := libgenders.IntersectionQuery{
				Left:  libgenders.ValueQuery{Expression: "attr4"},
				Right: libgenders.ValueQuery{Expression: "attr8=val8"},
			}

			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(libgenders.Set{0, 2}))
		})
	})

	context("DifferenceQuery", func() {
		it("returns a set matching the query", func() {
			query := libgenders.DifferenceQuery{
				Left:  libgenders.ValueQuery{Expression: "attr4"},
				Right: libgenders.ValueQuery{Expression: "attr8=val8"},
			}

			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(libgenders.Set{1, 3}))
		})
	})

	context("ComplementQuery", func() {
		it("returns a set matching the query", func() {
			query := libgenders.ComplementQuery{
				Query: libgenders.ValueQuery{Expression: "attr8=val8"},
			}

			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(libgenders.Set{1, 3, 5, 7}))
		})
	})
}
