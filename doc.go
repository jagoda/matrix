/*
Package matrix provides utilities for testing command line clients written in
Go. Testing command line clients in this way is typically done by including the
test cases in the main package so that the main() function can be called by the
test case.

	var (
		process *Process
	)

	BeforeEach(func() {
		process = CaptureProcess()
		defer process.Restore()

		main()
	})

	It("captures the output", func() {
		Expect(process.Output()).To(Equal("hello\n"))
	})
*/
package matrix
