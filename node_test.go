package libgenders_test

import (
	"testing"

	"github.com/ryanmoran/libgenders"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testNode(t *testing.T, context spec.G, it spec.S) {
	var Expect = NewWithT(t).Expect

	context("MergeAttributes", func() {
		it("merges the attributes of the given node into the existing node", func() {
			node := libgenders.NewNode("node1", map[string]string{
				"attr1": "",
				"attr2": "val2",
			})

			node.MergeAttributes(map[string]string{
				"attr3": "",
				"attr4": "val4",
			})

			Expect(node).To(Equal(libgenders.Node{
				Name: "node1",
				Attributes: map[string]string{
					"attr1": "",
					"attr2": "val2",
					"attr3": "",
					"attr4": "val4",
				},
			}))
		})
	})
}
