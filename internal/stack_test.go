package internal_test

import (
	"testing"

	"github.com/ryanmoran/libgenders/internal"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testStack(t *testing.T, _ spec.G, it spec.S) {
	var Expect = NewWithT(t).Expect

	it("acts like a stack for Tokens", func() {
		stack := internal.Stack{}
		stack.Push(internal.NewToken(internal.ValueTokenKind, "some-value"))
		stack.Push(internal.NewToken(internal.ValueTokenKind, "other-value"))
		Expect(stack.IsEmpty()).To(BeFalse())

		Expect(stack.Top()).To(Equal(internal.NewToken(internal.ValueTokenKind, "other-value")))
		Expect(stack.Pop()).To(Equal(internal.NewToken(internal.ValueTokenKind, "other-value")))
		Expect(stack.Pop()).To(Equal(internal.NewToken(internal.ValueTokenKind, "some-value")))
		Expect(stack.IsEmpty()).To(BeTrue())
	})
}
