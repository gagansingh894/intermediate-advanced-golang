package creational

import (
	"errors"
	"fmt"
)

type Vehicle interface {
	GetWheels() int
	GetSeats() int
}

type Car interface {
	GetDoors() int
}

type Motorbike interface {
	GetType() int
}

// ABSTRACT FACTORY
type VehicleFactory interface {
	GetVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized\n", f))

	}
}

// Car Factory

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct{}

func (c *CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}

type LuxuryCar struct{}

func (l *LuxuryCar) GetDoors() int {
	return 4
}

func (l *LuxuryCar) GetWheels() int {
	return 4
}

func (l *LuxuryCar) GetSeats() int {
	return 5
}

type FamilyCar struct{}

func (l *FamilyCar) GetDoors() int {
	return 5
}

func (l *FamilyCar) GetWheels() int {
	return 4
}

func (l *FamilyCar) GetSeats() int {
	return 5
}

// Motorbike Factory

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}

type SportMotorbike struct{}

func (s *SportMotorbike) GetWheels() int {
	return 2
}

func (s *SportMotorbike) GetSeats() int {
	return 1
}

func (s *SportMotorbike) GetType() int {
	return SportMotorbikeType
}

type CruiseMotorbike struct{}

func (c *CruiseMotorbike) GetWheels() int {
	return 2
}

func (c *CruiseMotorbike) GetSeats() int {
	return 2
}

func (c *CruiseMotorbike) GetType() int {
	return CruiseMotorbikeType
}

func AbstractFactoryExample() {
	carF, _ := GetVehicleFactory(CarFactoryType)
	fmt.Printf("%T\n", carF)
	luxuryCar, _ := carF.GetVehicle(LuxuryCarType)
	familyCar, _ := carF.GetVehicle(FamilyCarType)
	fmt.Printf("%T\n", luxuryCar)
	fmt.Printf("%T\n", familyCar)
}
