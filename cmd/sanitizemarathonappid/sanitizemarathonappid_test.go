package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("sanitizemarathonappid", func() {
	Describe("process", func() {
		It("works", func() {
			outputs := process([]string{"a.b.c.d.e", "a_b_c_d", "aBcDeF", "foo.λ___BAR"})
			Expect(outputs).To(Equal([]string{"a-b-c-d-e", "a-b-c-d", "abcdef", "foo-bar"}))
		})
	})

	Describe("main", func() {
		It("works", func() {
			main()
		})
	})
})
