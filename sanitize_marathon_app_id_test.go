package sanitizemarathonappid_test

import (
	. "github.com/msabramo/go-sanitize-marathon-app-id"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SanitizeMarathonAppId", func() {
	Describe("character substitution", func() {
		It("should replace a dot within text", func() {
			Expect(Sanitize("a.b")).To(Equal("a-b"))
		})

		It("should replace multiple dots", func() {
			Expect(Sanitize("a.b.c.d.e")).To(Equal("a-b-c-d-e"))
		})

		It("should replace an underscores within text", func() {
			Expect(Sanitize("a_b")).To(Equal("a-b"))
		})

		It("should replace multiple underscores", func() {
			Expect(Sanitize("a_b_c_d_e")).To(Equal("a-b-c-d-e"))
		})

		It("should work with any combination of different non-allowed characters", func() {
			Expect(Sanitize("a.b_c.d_e")).To(Equal("a-b-c-d-e"))
		})
	})

	Describe("uppercase letters", func() {
		It("should transform uppercase letters to lowercase letters", func() {
			Expect(Sanitize("aBcDe")).To(Equal("abcde"))
		})
	})

	Describe("consecutive non-allowed characters", func() {
		It("should not produce multiple consecutive hyphens", func() {
			Expect(Sanitize("a...b")).To(Equal("a-b"))
			Expect(Sanitize("a___b")).To(Equal("a-b"))
			Expect(Sanitize("a_._b")).To(Equal("a-b"))
		})
	})

	Describe("non-allowed, non-substitutable characters", func() {
		It("should remove any non-allowed character which is not substitutable", func() {
			Expect(Sanitize("fooλbar")).To(Equal("foobar"))
		})
	})

	It("should remove any separator at the beginning or end", func() {
		Expect(Sanitize(".foo_bar_")).To(Equal("foo-bar"))
	})

	It("should work with the combination of multiple cases", func() {
		Expect(Sanitize("foo.λ___BAR")).To(Equal("foo-bar"))
	})
})
