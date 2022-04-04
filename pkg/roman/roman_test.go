package roman_test

import (
	"example.com/tddworkshop/project/pkg/roman"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Roman", func() {
	Describe("TimeToRomanTime", func() {

		// Task: Extend the testsuite in `clock_test.go` to verify the functionality of

		var (
			clock = &roman.RomanClock{}
		)

		It("reports an error when invalid time format is used", func() {
			invalid := []string{"foo", "-5:00:00", ":00:00", "111:00:00", "00:01:02:03"}
			for _, i := range invalid {
				t, err := clock.TimeToRomanTime(i)
				Expect(err).To(HaveOccurred(), "expected '"+i+"' to raise an invalid format error")
				Expect(t).To(Equal(""))
			}
		})

		FIt("fails if hour is out of range", func() {
			_, err := clock.TimeToRomanTime("25:10:05")
			Expect(err).To(HaveOccurred())
		})

		PIt("converts 00:00")
		PIt("converts 05:08")
		PIt("converts 5:8")
		PIt("converts 11:22")
		PIt("converts 23:59")
	})

	Describe("CurrentRomanTime", func() {
		PIt("handles errors when getting the time from the ClockClient")
		PIt("returns the current time")
	})
})

// BONUS POINTS: coverage:
// ```
// ginkgo -cover; go tool cover -html=coverprofile.out
// ```
