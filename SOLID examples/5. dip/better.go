package dip

import "time"

type HttpService struct{}

func (s *HttpService) Alive() bool {
	return true
}

func (s *HttpService) Get(path string) (int32, string) {
	return 200, `What's the difference between Dubai and Abu Dhabi? ... 
		People in Dubai don't like the flintstones but people in Abu Dhabi doooo`
}

type SMSNotificationService struct{}

func (n *SMSNotificationService) Notify() {
	// do something to send an SMS
}

type MonitoredService interface {
	Alive() bool
}
type Notifier interface {
	Notify()
}

type Watchdog struct{}

func (Watchdog) Watch(s MonitoredService, n Notifier) {
	for {
		if !s.Alive() {
			n.Notify()
		}

		time.Sleep(time.Second)
	}
}
