package root_test

import (
	"github.com/egoholic/spec/node"
	. "github.com/egoholic/spec/root"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	name1        = "Node1"
	description1 = "Description1"
	root1        = New(name1, description1)
)

var _ = Describe("Root package", func() {
	Describe("Root", func() {
		Describe("New()", func() {
			It("returns Root", func() {
				root := New(name1, description1)
				Expect(root).To(BeAssignableToTypeOf(&Root{}))
				Expect(root.Name()).To(Equal(name1))
				Expect(root.Description()).To(Equal(description1))
			})
		})

		Describe(".GetNode()", func() {
			It("returns Root", func() {
				Expect(root1.GetNode()).To(BeAssignableToTypeOf(&node.Node{}))
			})
		})

		Describe(".Name()", func() {
			It("returns Root's name", func() {
				Expect(root1.Name()).To(Equal(name1))
			})
		})

		Describe(".Description()", func() {
			It("returns Root's name", func() {
				Expect(root1.Description()).To(Equal(description1))
			})
		})

		Describe(".Validate()", func() {
			Context("when valid", func() {
				It("returns Root's name", func() {
					_, ok := root1.Validate(nil)
					Expect(ok).To(BeTrue())
				})
			})
		})
	})
})
