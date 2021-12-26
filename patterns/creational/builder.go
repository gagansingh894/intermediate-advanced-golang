package creational

import "fmt"

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// Director
type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// Product
type VehicleProduct struct {
	Wheels int
	Seats int
	Structure string
}

// Car type builder
type CarBuilder struct {
	v	VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// Bike type builder
type BikeBuilder struct {
	v	VehicleProduct
}

func (c *BikeBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}

func (c *BikeBuilder) SetSeats() BuildProcess {
	c.v.Seats = 2
	return c
}

func (c *BikeBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Motorbike"
	return c
}

func (c *BikeBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// Bus type builder
type BusBuilder struct {
	v	VehicleProduct
}

func (c *BusBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 8
	return c
}

func (c *BusBuilder) SetSeats() BuildProcess {
	c.v.Seats = 16
	return c
}

func (c *BusBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Bus"
	return c
}

func (c *BusBuilder) GetVehicle() VehicleProduct {
	return c.v
}

func BuilderExample() {
	carBuilder := &CarBuilder{}
	bikeBuilder := &BikeBuilder{}
	busBuilder := &BusBuilder{}

	director := ManufacturingDirector{}

	// Build car
	director.SetBuilder(carBuilder)
	director.Construct()
	car := director.builder.GetVehicle()
	fmt.Printf(`
	Building car....
	Wheels: %d
	Seats: %d
	Structure: %s
	`, car.Wheels, car.Seats, car.Structure)

	// Build bike
	director.SetBuilder(bikeBuilder)
	director.Construct()
	bike := director.builder.GetVehicle()
	fmt.Printf(`
	Building bike....
	Wheels: %d
	Seats: %d
	Structure: %s
	`, bike.Wheels, bike.Seats, bike.Structure)

	// Build bus
	director.SetBuilder(busBuilder)
	director.Construct()
	bus := director.builder.GetVehicle()
	fmt.Printf(`
	Building bus....
	Wheels: %d
	Seats: %d
	Structure: %s
	`, bus.Wheels, bus.Seats, bus.Structure)
}
