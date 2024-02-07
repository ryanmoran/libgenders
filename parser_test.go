package libgenders_test

import (
	"testing"

	"github.com/ryanmoran/libgenders"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testParser(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
		parser libgenders.Parser
	)

	it.Before(func() {
		parser = libgenders.NewParser()
	})

	context("Parse", func() {
		it("parses a line into a list of nodes", func() {
			nodes, err := parser.Parse("node1 attr1,attr2=val2")
			Expect(err).NotTo(HaveOccurred())
			Expect(nodes).To(Equal([]libgenders.Node{
				{
					Name: "node1",
					Attributes: map[string]string{
						"attr1": "",
						"attr2": "val2",
					},
				},
			}))
		})

		context("when there are multiple node names on a line", func() {
			it("parses a line into a list of nodes", func() {
				nodes, err := parser.Parse("node1,node2 attr1,attr2=val2")
				Expect(err).NotTo(HaveOccurred())
				Expect(nodes).To(Equal([]libgenders.Node{
					{
						Name: "node1",
						Attributes: map[string]string{
							"attr1": "",
							"attr2": "val2",
						},
					},
					{
						Name: "node2",
						Attributes: map[string]string{
							"attr1": "",
							"attr2": "val2",
						},
					},
				}))
			})
		})

		context("when the line begins with a comment", func() {
			it("returns an empty list of nodes", func() {
				nodes, err := parser.Parse("# comment")
				Expect(err).NotTo(HaveOccurred())
				Expect(nodes).To(BeEmpty())
			})
		})

		context("when the line ends with a comment", func() {
			it("returns an empty list of nodes", func() {
				nodes, err := parser.Parse("node1,node2 attr1,attr2=val2# comment")
				Expect(err).NotTo(HaveOccurred())
				Expect(nodes).To(Equal([]libgenders.Node{
					{
						Name: "node1",
						Attributes: map[string]string{
							"attr1": "",
							"attr2": "val2",
						},
					},
					{
						Name: "node2",
						Attributes: map[string]string{
							"attr1": "",
							"attr2": "val2",
						},
					},
				}))
			})
		})

		context("when the line is whitespace only", func() {
			it("returns an empty list of nodes", func() {
				nodes, err := parser.Parse("                 ")
				Expect(err).NotTo(HaveOccurred())
				Expect(nodes).To(BeEmpty())
			})
		})

		context("when the node name includes a range", func() {
			it("expands the range", func() {
				nodes, err := parser.Parse("node[1-2]name attr1,attr2=val2")
				Expect(err).NotTo(HaveOccurred())
				Expect(nodes).To(Equal([]libgenders.Node{
					{
						Name: "node1name",
						Attributes: map[string]string{
							"attr1": "",
							"attr2": "val2",
						},
					},
					{
						Name: "node2name",
						Attributes: map[string]string{
							"attr1": "",
							"attr2": "val2",
						},
					},
				}))
			})

			context("when the range is a single value", func() {
				it("expands the range", func() {
					nodes, err := parser.Parse("node[1] attr1,attr2=val2")
					Expect(err).NotTo(HaveOccurred())
					Expect(nodes).To(Equal([]libgenders.Node{
						{
							Name: "node1",
							Attributes: map[string]string{
								"attr1": "",
								"attr2": "val2",
							},
						},
					}))
				})
			})

			context("when the range is a comma-separated list", func() {
				it("expands the range", func() {
					nodes, err := parser.Parse("node[1,2] attr1,attr2=val2")
					Expect(err).NotTo(HaveOccurred())
					Expect(nodes).To(Equal([]libgenders.Node{
						{
							Name: "node1",
							Attributes: map[string]string{
								"attr1": "",
								"attr2": "val2",
							},
						},
						{
							Name: "node2",
							Attributes: map[string]string{
								"attr1": "",
								"attr2": "val2",
							},
						},
					}))
				})
			})
		})

		context("when the line only specifies a node name", func() {
			it("expands the range", func() {
				nodes, err := parser.Parse("node1")
				Expect(err).NotTo(HaveOccurred())
				Expect(nodes).To(Equal([]libgenders.Node{
					{Name: "node1"},
				}))
			})

			context("when the range is a single value", func() {
				it("expands the range", func() {
					nodes, err := parser.Parse("node[1] attr1,attr2=val2")
					Expect(err).NotTo(HaveOccurred())
					Expect(nodes).To(Equal([]libgenders.Node{
						{
							Name: "node1",
							Attributes: map[string]string{
								"attr1": "",
								"attr2": "val2",
							},
						},
					}))
				})
			})
		})
	})
}
