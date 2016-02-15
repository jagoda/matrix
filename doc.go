/*
Package matrix provides utilities for testing command line clients written in
Go. Testing command line clients in this way is typically done by including the
test cases in the main package so that the main() function can be called by the
test case.

	package main

	import (
		. "github.com/onsi/ginkgo"
		. "github.com/onsi/gomega"

		"github.com/jagoda/matrix"
	)

	var _ = Describe("Process", func() {
		var (
			process *matrix.Process
		)

		BeforeEach(func() {
			process = matrix.CaptureProcess()
			defer process.Restore()

			main()
		})

		It("captures the output", func() {
			Expect(process.Output()).To(Equal("hello\n"))
		})
	})
*/
package matrix
