package search_test

import (
	"testing"

	"github.com/ryanmoran/libgenders/search"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testQuery(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		attrs = map[string]search.Set{
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
		attrvals = map[string]search.Set{
			"attr10=val10": {1, 3, 5, 7},
			"attr2=val2":   {0, 1, 2, 3, 4, 5, 6, 7},
			"attr4=val4":   {0, 1, 2, 3},
			"attr6=val6":   {4, 5, 6, 7},
			"attr8=val8":   {0, 2, 4, 6},
		}
		indices = search.Set{0, 1, 2, 3, 4, 5, 6, 7}
	)

	context("ParseQuery", func() {
		it("parses attribute-only queries", func() {
			tokens := []search.Token{
				{Kind: search.ValueTokenKind, Text: "attr1"},
			}

			Expect(search.ParseQuery(tokens)).To(Equal(search.ValueQuery{
				Expression: "attr1",
			}))
		})

		it("parses attribute-value queries", func() {
			tokens := []search.Token{
				{Kind: search.ValueTokenKind, Text: "attr2=val2"},
			}

			Expect(search.ParseQuery(tokens)).To(Equal(search.ValueQuery{
				Expression: "attr2=val2",
			}))
		})

		it("parses union queries", func() {
			tokens := []search.Token{
				{Kind: search.UnionTokenKind, Text: "||"},
				{Kind: search.ValueTokenKind, Text: "attr2"},
				{Kind: search.ValueTokenKind, Text: "attr1"},
			}

			Expect(search.ParseQuery(tokens)).To(Equal(search.UnionQuery{
				Left:  search.ValueQuery{Expression: "attr1"},
				Right: search.ValueQuery{Expression: "attr2"},
			}))
		})

		it("parses intersection queries", func() {
			tokens := []search.Token{
				{Kind: search.IntersectionTokenKind, Text: "&&"},
				{Kind: search.ValueTokenKind, Text: "attr2"},
				{Kind: search.ValueTokenKind, Text: "attr1"},
			}

			Expect(search.ParseQuery(tokens)).To(Equal(search.IntersectionQuery{
				Left:  search.ValueQuery{Expression: "attr1"},
				Right: search.ValueQuery{Expression: "attr2"},
			}))
		})

		it("parses difference queries", func() {
			tokens := []search.Token{
				{Kind: search.DifferenceTokenKind, Text: "--"},
				{Kind: search.ValueTokenKind, Text: "attr2"},
				{Kind: search.ValueTokenKind, Text: "attr1"},
			}

			Expect(search.ParseQuery(tokens)).To(Equal(search.DifferenceQuery{
				Left:  search.ValueQuery{Expression: "attr1"},
				Right: search.ValueQuery{Expression: "attr2"},
			}))
		})

		it("parses complement queries", func() {
			tokens := []search.Token{
				{Kind: search.ComplementTokenKind, Text: "~"},
				{Kind: search.ValueTokenKind, Text: "attr1"},
			}

			Expect(search.ParseQuery(tokens)).To(Equal(search.ComplementQuery{
				Query: search.ValueQuery{Expression: "attr1"},
			}))
		})

		it("parses mixed queries", func() {
			tokens := []search.Token{
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
			}

			Expect(search.ParseQuery(tokens)).To(Equal(search.IntersectionQuery{
				Left: search.UnionQuery{
					Left: search.IntersectionQuery{
						Left: search.ValueQuery{Expression: "attr1"},
						Right: search.ComplementQuery{
							Query: search.ValueQuery{Expression: "attr3"},
						},
					},
					Right: search.DifferenceQuery{
						Left:  search.ValueQuery{Expression: "attr1"},
						Right: search.ValueQuery{Expression: "attr5"},
					},
				},
				Right: search.ValueQuery{Expression: "attr7"},
			}))
		})
	})

	context("ValueQuery", func() {
		it("returns a set matching the attribute-only query", func() {
			query := search.ValueQuery{Expression: "attr1"}
			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(search.Set{0, 1, 2, 3, 4, 5, 6, 7}))
		})

		it("returns a set matching the attribute-value query", func() {
			query := search.ValueQuery{Expression: "attr4=val4"}
			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(search.Set{0, 1, 2, 3}))
		})
	})

	context("UnionQuery", func() {
		it("returns a set matching the query", func() {
			query := search.UnionQuery{
				Left:  search.ValueQuery{Expression: "attr4"},
				Right: search.ValueQuery{Expression: "attr8=val8"},
			}

			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(search.Set{0, 1, 2, 3, 4, 6}))
		})
	})

	context("IntersectionQuery", func() {
		it("returns a set matching the query", func() {
			query := search.IntersectionQuery{
				Left:  search.ValueQuery{Expression: "attr4"},
				Right: search.ValueQuery{Expression: "attr8=val8"},
			}

			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(search.Set{0, 2}))
		})
	})

	context("DifferenceQuery", func() {
		it("returns a set matching the query", func() {
			query := search.DifferenceQuery{
				Left:  search.ValueQuery{Expression: "attr4"},
				Right: search.ValueQuery{Expression: "attr8=val8"},
			}

			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(search.Set{1, 3}))
		})
	})

	context("ComplementQuery", func() {
		it("returns a set matching the query", func() {
			query := search.ComplementQuery{
				Query: search.ValueQuery{Expression: "attr8=val8"},
			}

			result := query.Evaluate(attrs, attrvals, indices)
			Expect(result).To(Equal(search.Set{1, 3, 5, 7}))
		})
	})
}
