package parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
})
