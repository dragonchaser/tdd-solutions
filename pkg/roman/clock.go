package roman

import (
	"errors"
	"regexp"
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

// CurrentRomanTime() returns the time in roman numbers, e.g. XII:XV:IIX (12:15:08)
func (*RomanClock) CurrentRomanTime() (string, error) {
	return "", nil
}

// TimeToRomanTime() transcribes a time string into roman numbers, , e.g. XII:XV:IIX (12:15:08)
func (*RomanClock) TimeToRomanTime(time string) (string, error) {
	// POSIBLE SOLUTION FOR FIRST ITERATION:
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
	if len(matches) != 1 {
		return "", errors.New("Invalid clock format")
	}
	return "", nil
}
