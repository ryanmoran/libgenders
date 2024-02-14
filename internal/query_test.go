package internal_test

import (
	"testing"

	"github.com/ryanmoran/libgenders/internal"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testQuery(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		attrs = map[string]internal.Set{
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
		attrvals = map[string]internal.Set{
			"attr10=val10": {1, 3, 5, 7},
			"attr2=val2":   {0, 1, 2, 3, 4, 5, 6, 7},
			"attr4=val4":   {0, 1, 2, 3},
			"attr6=val6":   {4, 5, 6, 7},
			"attr8=val8":   {0, 2, 4, 6},
		}
		indices = internal.Set{0, 1, 2, 3, 4, 5, 6, 7}
	)

	context("ParseQuery", func() {
		it("parses attribute-only queries", func() {
			tokens := []internal.Token{
				{Kind: internal.ValueTokenKind, Text: "attr1"},
			}

			Expect(internal.ParseQuery(tokens)).To(Equal(internal.ValueQuery{
				Expression: "attr1",
			}))
		})

		it("parses attribute-value queries", func() {
			tokens := []internal.Token{
				{Kind: internal.ValueTokenKind, Text: "attr2=val2"},
			}

			Expect(internal.ParseQuery(tokens)).To(Equal(internal.ValueQuery{
				Expression: "attr2=val2",
			}))
		})

		it("parses union queries", func() {
			tokens := []internal.Token{
				{Kind: internal.UnionTokenKind, Text: "||"},
				{Kind: internal.ValueTokenKind, Text: "attr2"},
				{Kind: internal.ValueTokenKind, Text: "attr1"},
			}

			Expect(internal.ParseQuery(tokens)).To(Equal(internal.UnionQuery{
				Left:  internal.ValueQuery{Expression: "attr1"},
				Right: internal.ValueQuery{Expression: "attr2"},
			}))
		})

		it("parses intersection queries", func() {
			tokens := []internal.Token{
				{Kind: internal.IntersectionTokenKind, Text: "&&"},
				{Kind: internal.ValueTokenKind, Text: "attr2"},
				{Kind: internal.ValueTokenKind, Text: "attr1"},
			}

			Expect(internal.ParseQuery(tokens)).To(Equal(internal.IntersectionQuery{
				Left:  internal.ValueQuery{Expression: "attr1"},
				Right: internal.ValueQuery{Expression: "attr2"},
			}))
		})

		it("parses difference queries", func() {
			tokens := []internal.Token{
				{Kind: internal.DifferenceTokenKind, Text: "--"},
				{Kind: internal.ValueTokenKind, Text: "attr2"},
				{Kind: internal.ValueTokenKind, Text: "attr1"},
			}

			Expect(internal.ParseQuery(tokens)).To(Equal(internal.DifferenceQuery{
				Left:  internal.ValueQuery{Expression: "attr1"},
				Right: internal.ValueQuery{Expression: "attr2"},
			}))
		})

		it("parses complement queries", func() {
			tokens := []internal.Token{
				{Kind: internal.ComplementTokenKind, Text: "~"},
				{Kind: internal.ValueTokenKind, Text: "attr1"},
			}

			Expect(internal.ParseQuery(tokens)).To(Equal(internal.ComplementQuery{
				Query: internal.ValueQuery{Expression: "attr1"},
			}))
		})

		it("parses mixed queries", func() {
			tokens := []internal.Token{
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
			}

			Expect(internal.ParseQuery(tokens)).To(Equal(internal.IntersectionQuery{
				Left: internal.UnionQuery{
					Left: internal.IntersectionQuery{
						Left: internal.ValueQuery{Expression: "attr1"},
						Right: internal.ComplementQuery{
							Query: internal.ValueQuery{Expression: "attr3"},
						},
					},
					Right: internal.DifferenceQuery{
						Left:  internal.ValueQuery{Expression: "attr1"},
						Right: internal.ValueQuery{Expression: "attr5"},
					},
				},
				Right: internal.ValueQuery{Expression: "attr7"},
			}))
		})
	})

	context("ValueQuery", func() {
		it("returns a set matching the attribute-only query", func() {
			query := internal.ValueQuery{Expression: "attr1"}
			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(internal.Set{0, 1, 2, 3, 4, 5, 6, 7}))
		})

		it("returns a set matching the attribute-value query", func() {
			query := internal.ValueQuery{Expression: "attr4=val4"}
			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(internal.Set{0, 1, 2, 3}))
		})
	})

	context("UnionQuery", func() {
		it("returns a set matching the query", func() {
			query := internal.UnionQuery{
				Left:  internal.ValueQuery{Expression: "attr4"},
				Right: internal.ValueQuery{Expression: "attr8=val8"},
			}

			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(internal.Set{0, 1, 2, 3, 4, 6}))
		})
	})

	context("IntersectionQuery", func() {
		it("returns a set matching the query", func() {
			query := internal.IntersectionQuery{
				Left:  internal.ValueQuery{Expression: "attr4"},
				Right: internal.ValueQuery{Expression: "attr8=val8"},
			}

			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(internal.Set{0, 2}))
		})
	})

	context("DifferenceQuery", func() {
		it("returns a set matching the query", func() {
			query := internal.DifferenceQuery{
				Left:  internal.ValueQuery{Expression: "attr4"},
				Right: internal.ValueQuery{Expression: "attr8=val8"},
			}

			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(internal.Set{1, 3}))
		})
	})

	context("ComplementQuery", func() {
		it("returns a set matching the query", func() {
			query := internal.ComplementQuery{
				Query: internal.ValueQuery{Expression: "attr8=val8"},
			}

			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(internal.Set{1, 3, 5, 7}))
		})
	})
}
