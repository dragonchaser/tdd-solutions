package isp

type Chargeable interface {
	ChargeTo(int8)
}

type Pumpable interface {
	PumpUp(float32)
}

type Charger struct{}

func (charger *Charger) Charge(c Chargeable) {
	c.ChargeTo(100)
}

type Pump struct{}

func (pump *Pump) Pump(s Pumpable) {
	s.PumpUp(3)
}

type Bike struct {
	pressure float32
}

func (b *Bike) PumpUp(p float32) {
	b.pressure = p
}

type Ebike struct {
	Bike

	battery int8
}

func (e *Ebike) ChargeTo(b int8) {
	e.battery = b
}
