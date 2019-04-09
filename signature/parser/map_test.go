package parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TODO: implement tests
var _ = Describe("Parsing", func() {
	Describe("ParseMap()", func() {
		Context("when map sigrature given", func() {
			It("returns MapNode", func() {
				Expect(true).To(Equal(true))
			})
		})

		Context("when not map signature given", func() {
			It("returns error", func() {
				Expect(true).To(Equal(true))
			})
		})
	})

	Describe("*MapSignature", func() {
		Describe(".GolangSignature()", func() {
			It("returns golang signature", func() {
				Expect(true).To(Equal(true))
			})
		})

		Describe(".Matches()", func() {
			Context("when matches", func() {
				It("returns true", func() {
					Expect(true).To(Equal(true))

				})
			})

			Context("when does not matche", func() {
				It("returns false", func() {
					Expect(true).To(Equal(true))
				})
			})
		})
	})
})
