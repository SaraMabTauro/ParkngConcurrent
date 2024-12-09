package simulation

import (
	"fmt"
	"math/rand"
	"time"
)

type Simulator struct {
	ParkingLot   *ParkingLot
	VehicleCount int
}

func NewSimulator(capacity int) *Simulator {
	return &Simulator{
		ParkingLot: NewParkingLot(capacity),
	}
}

func (s *Simulator) Run() {
	for i := 0; i < s.VehicleCount; i++ {
		vehicle := NewVehicle(i)

		go func(v *Vehicle) {
			for !s.ParkingLot.EnterVehicle(v) {
				time.Sleep(time.Second)
			}

			fmt.Printf("Vehículo %d estacionado por %v.\n", v.ID, v.ParkingTime)
			v.StayParked()
			s.ParkingLot.ExitVehicle(v)
			fmt.Printf("Vehículo %d ha salido del estacionamiento.\n", v.ID)
		}(vehicle)

		// tiempo de llegada de nuevos vehículos
		time.Sleep(time.Duration(500+rand.Intn(500)) * time.Millisecond)
	}
}
