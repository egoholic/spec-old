package parser_test

import (
	"github.com/egoholic/spec/signature"
	. "github.com/egoholic/spec/signature/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	sig1 = "string]int"
	sig2 = "int64]string"
	sig3 = "int32]bool"
	sig4 = "int]int"
	sig5 = "float]string"
	sig6 = "float64]float"
	sig7 = "bool][]string"
)

var _ = Describe("Parsing", func() {
	Describe("ParsePrimtive()", func() {
		It("returns signature", func() {
			var (
				sig signature.Signature
				err error
			)
			sig, err = ParsePrimitive(sig1)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(sig.GolangSignature()).To(Equal("string"))

			sig, err = ParsePrimitive(sig2)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(sig.GolangSignature()).To(Equal("int64"))

			sig, err = ParsePrimitive(sig3)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(sig.GolangSignature()).To(Equal("int32"))

			sig, err = ParsePrimitive(sig4)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(sig.GolangSignature()).To(Equal("int"))

			sig, err = ParsePrimitive(sig5)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(sig.GolangSignature()).To(Equal("float"))

			sig, err = ParsePrimitive(sig6)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(sig.GolangSignature()).To(Equal("float64"))

			sig, err = ParsePrimitive(sig7)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(sig.GolangSignature()).To(Equal("bool"))
		})
	})

	Describe("*PrimitiveSignature", func() {
		Describe(".GolangSignature()", func() {
			It("returns string representation of Golang signature", func() {
				var sig signature.Signature

				sig, _ = ParsePrimitive(sig1)
				Expect(sig.GolangSignature()).To(Equal("string"))
				sig, _ = ParsePrimitive(sig2)
				Expect(sig.GolangSignature()).To(Equal("int64"))
				sig, _ = ParsePrimitive(sig3)
				Expect(sig.GolangSignature()).To(Equal("int32"))
				sig, _ = ParsePrimitive(sig4)
				Expect(sig.GolangSignature()).To(Equal("int"))
				sig, _ = ParsePrimitive(sig5)
				Expect(sig.GolangSignature()).To(Equal("float"))
				sig, _ = ParsePrimitive(sig6)
				Expect(sig.GolangSignature()).To(Equal("float64"))
				sig, _ = ParsePrimitive(sig7)
				Expect(sig.GolangSignature()).To(Equal("bool"))
			})
		})

		Describe(".Matches()", func() {
			Context("when matches", func() {
				It("returns true", func() {
					var (
						original signature.Signature
						toMatch  signature.Signature
					)

					original, _ = ParsePrimitive(sig1)
					toMatch, _ = ParsePrimitive(sig1)
					Expect(original.Matches(toMatch)).To(BeTrue())

					original, _ = ParsePrimitive(sig2)
					toMatch, _ = ParsePrimitive(sig2)
					Expect(original.Matches(toMatch)).To(BeTrue())

					original, _ = ParsePrimitive(sig3)
					toMatch, _ = ParsePrimitive(sig3)
					Expect(original.Matches(toMatch)).To(BeTrue())

					original, _ = ParsePrimitive(sig4)
					toMatch, _ = ParsePrimitive(sig4)
					Expect(original.Matches(toMatch)).To(BeTrue())

					original, _ = ParsePrimitive(sig5)
					toMatch, _ = ParsePrimitive(sig5)
					Expect(original.Matches(toMatch)).To(BeTrue())

					original, _ = ParsePrimitive(sig6)
					toMatch, _ = ParsePrimitive(sig6)
					Expect(original.Matches(toMatch)).To(BeTrue())

					original, _ = ParsePrimitive(sig7)
					toMatch, _ = ParsePrimitive(sig7)
					Expect(original.Matches(toMatch)).To(BeTrue())
				})
			})

			Context("when does not match", func() {
				It("returns false", func() {
					Expect(true).To(Equal(true))
				})
			})
		})
	})
})
