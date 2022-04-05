package roman

import (
	"errors"
	"regexp"
	"strconv"
)

//go:generate mockery --name ClockClient

type ClockClient interface {
	// returns time string in format: `hh:mm:ss`
	// this would invoke an hypothetical api call
	CurrentTime() (string, error)
}

type RomanClock struct {
	clockClient ClockClient
}

func New(cc ClockClient) *RomanClock {
	return &RomanClock{clockClient: cc}
}

// CurrentRomanTime() returns the time in roman numbers, e.g. XII:XV:IIX (12:15:08)
func (rc *RomanClock) CurrentRomanTime() (string, error) {
	localRomanTime, err := rc.clockClient.CurrentTime()
	if err != nil {
		return "", err
	}
	return rc.TimeToRomanTime(localRomanTime)
}

// TimeToRomanTime() transcribes a time string into roman numbers, , e.g. XII:XV:IIX (12:15:08)
func (rc *RomanClock) TimeToRomanTime(time string) (string, error) {
	// POSIBLE SOLUTION FOR FIRST ITERATION:
	//
	//matched, err := regexp.Match(`^\d{1,2}:\d{1,2}:\d{1,2}$`, []byte(time))
	//if err != nil {
	//	return "", err
	//}
	//if !matched {
	//	return "", errors.New("Invalid clock format")
	//}
	//return "", nil

	r := regexp.MustCompile(`^(\d{1,2}):(\d{1,2}):(\d{1,2})$`)
	matches := r.FindStringSubmatch(time)
	if len(matches) != 4 { // [0]: matched string, [1-3]: capture groups
		return "", errors.New("invalid time format")
	}
	hours, err := strconv.ParseInt(matches[1], 10, 64)
	if hours > 23 || err != nil {
		return "", errors.New("invalid value for hour")
	}

	minutes, err := strconv.ParseInt(matches[2], 10, 64)
	if minutes > 59 || err != nil {
		return "", errors.New("invalid value for minute")
	}

	seconds, err := strconv.ParseInt(matches[3], 10, 64)
	if seconds > 59 || err != nil {
		return "", errors.New("invalid value for minute")
	}

	return rc.intToRoman(int(hours)) + ":" + rc.intToRoman(int(minutes)) + ":" + rc.intToRoman(int(seconds)), nil
}

func (rc *RomanClock) intToRoman(num int) string {
	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}

	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I"}
	roman := ""
	i := 0
	for num > 0 {
		k := num / values[i]
		for j := 0; j < k; j++ {
			roman += symbols[i]
			num -= values[i]
		}
		i++
	}
	return roman
}
