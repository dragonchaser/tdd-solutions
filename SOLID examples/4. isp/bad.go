package isp

type ChargeableVehicle interface {
	ChargeTo(percentage int8)
	PumpUp(pressure float32)
}

type Charger struct{}

func (c *Charger) Charge(b ChargeableVehicle) {
	b.ChargeTo(100)
}

type Pump struct{}

func (p *Pump) Pump(b ChargeableVehicle) {
	b.PumpUp(3)
}

type Ebike struct {
	battery  int8
	pressure float32
}

func (e *Ebike) ChargeTo(b int8) {
	e.battery = b
}

func (e *Ebike) PumpUp(p float32) {
	e.pressure = p
}

type RoadBike struct {
	pressure float32
}

func (b *RoadBike) PumpUp(p float32) {
	b.pressure = p
}

func (b *RoadBike) Charge(_ float32) {
	// nop. meh
}
