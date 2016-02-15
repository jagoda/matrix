package matrix_test

import (
	. "github.com/jagoda/matrix"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
)

var _ = Describe("Process", func() {
	Describe("Capturing process output", func() {
		var (
			process *Process
		)

		BeforeEach(func() {
			process = CaptureProcess()
			defer process.Restore()

			os.Stdout.WriteString("this is standard output\n")
			os.Stderr.WriteString("this is standard error\n")
			os.Stdout.WriteString("this is standard output again\n")
		})

		It("captures the combined output", func() {
			Expect(process.Output()).To(Equal(
				"this is standard output\n" +
					"this is standard error\n" +
					"this is standard output again\n",
			))
		})
	})

	Describe("Managing process arguements", func() {
		Context("when no arguments are specified", func() {
			var (
				before, during []string
			)

			BeforeEach(func() {
				before = os.Args

				process := CaptureProcess()
				defer process.Restore()

				during = os.Args
			})

			It("masks the process arguments", func() {
				Expect(during).NotTo(Equal(before))
				Expect(during).To(Equal([]string{"test"}))
			})

			It("restores the process arguments", func() {
				Expect(os.Args).To(Equal(before))
			})
		})

		Context("when arguments are specified", func() {
			var (
				before, during []string
			)

			BeforeEach(func() {
				before = os.Args

				process := CaptureProcess()
				defer process.Restore()

				process.SetArgs("foo", "bar", "baz")
				during = os.Args
			})

			It("updates the process arguments", func() {
				Expect(during).NotTo(Equal(before))
				Expect(during).To(Equal([]string{"foo", "bar", "baz"}))
			})

			It("restores the process arguments", func() {
				Expect(os.Args).To(Equal(before))
			})
		})
	})
})
