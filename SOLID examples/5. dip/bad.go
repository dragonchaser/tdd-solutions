package dip

import "time"

type HttpService struct{}

func (s *HttpService) Get(path string) (int32, string) {
	return 200, "Why did the bicycle fall over? ... It was two-tired."
}

type SMSNotificationService struct{}

func (n *SMSNotificationService) SendAlertSMS() {
	// do something to send an SMS
}

type Watchdog struct{}

func (Watchdog) Watch(s *HttpService, n *SMSNotificationService) {
	for {
		code, _ := s.Get("/")
		if code != 200 {
			n.SendAlertSMS()
		}

		time.Sleep(time.Second)
	}
}
