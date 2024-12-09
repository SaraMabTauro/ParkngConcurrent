package domain

import (
	"math/rand"
	"time"
)

type Vehicle struct {
	ID          int
	ParkingTime time.Duration
	SpaceID     int
}

func NewVehicle(id int) *Vehicle {
	minTime := 3
	maxTime := 5
	parkingTime := time.Duration(minTime+rand.Intn(maxTime-minTime)) * time.Second

	return &Vehicle{
		ID:          id,
		ParkingTime: parkingTime,
		SpaceID:     -1,
	}
}

func (v *Vehicle) StayParked() {
	time.Sleep(v.ParkingTime)
}

func (v *Vehicle) AssignParkingSpace(spaceID int) {
	v.SpaceID = spaceID
}

func (v *Vehicle) RemoveParkingSpace() {
	v.SpaceID = -1
}

func (v *Vehicle) GetParkingSpace() int {
	return v.SpaceID
}

func (v *Vehicle) IsParked() bool {
	return v.SpaceID != -1
}
