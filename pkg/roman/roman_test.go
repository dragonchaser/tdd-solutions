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

		It("fails if hour is out of range", func() {
			_, err := clock.TimeToRomanTime("25:10:05")
			Expect(err).To(HaveOccurred())
		})

		It("fails if minute is out of range", func() {
			_, err := clock.TimeToRomanTime("25:60:05")
			Expect(err).To(HaveOccurred())
		})

		It("fails if second is out of range", func() {
			_, err := clock.TimeToRomanTime("25:10:60")
			Expect(err).To(HaveOccurred())
		})

		It("converts 00:00:00", func() {
			time, err := clock.TimeToRomanTime("00:00:00")
			Expect(err).ToNot(HaveOccurred())
			Expect(time).To(Equal("::"))
		})

		It("converts 05:08:25", func() {
			time, err := clock.TimeToRomanTime("05:08:25")
			Expect(err).ToNot(HaveOccurred())
			Expect(time).To(Equal("V:VIII:XXV"))
		})

		It("converts 5:8:25", func() {
			time, err := clock.TimeToRomanTime("05:08:25")
			Expect(err).ToNot(HaveOccurred())
			Expect(time).To(Equal("V:VIII:XXV"))
		})
		It("converts 11:22:59", func() {
			time, err := clock.TimeToRomanTime("11:22:59")
			Expect(err).ToNot(HaveOccurred())
			Expect(time).To(Equal("XI:XXII:LIX"))
		})
		It("converts 23:59:11", func() {
			time, err := clock.TimeToRomanTime("23:59:11")
			Expect(err).ToNot(HaveOccurred())
			Expect(time).To(Equal("XXIII:LIX:XI"))
		})
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
