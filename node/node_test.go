package node_test

import (
	. "github.com/egoholic/spec/node"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	name1        = "Node1"
	description1 = "Description1"
	node1        = New(name1, description1)
)

var _ = Describe("Node package", func() {
	Describe("Node", func() {
		Describe("New()", func() {
			It("returns Node", func() {
				node := New(name1, description1)
				Expect(node).To(BeAssignableToTypeOf(&Node{}))
				Expect(node.Name()).To(Equal(name1))
				Expect(node.Description()).To(Equal(description1))
			})
		})

		Describe(".GetNode()", func() {
			It("returns Node", func() {
				Expect(node1.GetNode()).To(Equal(node1))
			})
		})

		Describe(".Name()", func() {
			It("returns Node's name", func() {
				Expect(node1.Name()).To(Equal(name1))
			})
		})

		Describe(".Description()", func() {
			It("returns Node's name", func() {
				Expect(node1.Description()).To(Equal(description1))
			})
		})

		Describe(".Validate()", func() {
			Context("when valid", func() {
				It("returns Node's name", func() {
					_, ok := node1.Validate(nil)
					Expect(ok).To(BeTrue())
				})
			})
		})
	})
})
