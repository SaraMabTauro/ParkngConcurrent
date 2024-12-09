package simulation

import (
	"math/rand"
	"time"
)

type Vehicle struct {
	ID          int
	ParkingTime time.Duration
}

func NewVehicle(id int) *Vehicle {
	// tiempo de estacionamiento entre 3 y 5 segundos
	parkingTime := time.Duration(3+rand.Intn(3)) * time.Second
	return &Vehicle{
		ID:          id,
		ParkingTime: parkingTime,
	}
}

func (v *Vehicle) StayParked() {
	time.Sleep(v.ParkingTime)
}
