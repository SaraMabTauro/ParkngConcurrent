package simulation

import (
	"fmt"
)

type ParkingLot struct {
	Capacity       int
	CurrentCount   int
	EntrySemaphore *Semaphore
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		Capacity:       capacity,
		EntrySemaphore: NewSemaphore(1), // Un vehículo a la vez en entrada/salida
	}
}

func (p *ParkingLot) EnterVehicle(vehicle *Vehicle) bool {
	if p.CurrentCount >= p.Capacity {
		fmt.Printf("Estacionamiento lleno, el vehículo %d está esperando.\n", vehicle.ID)
		return false
	}

	p.EntrySemaphore.Acquire() // Bloquea la entrada/salida
	defer p.EntrySemaphore.Release()

	fmt.Printf("Vehículo %d está entrando al estacionamiento.\n", vehicle.ID)
	p.CurrentCount++
	return true
}

func (p *ParkingLot) ExitVehicle(vehicle *Vehicle) {
	p.EntrySemaphore.Acquire()
	defer p.EntrySemaphore.Release()

	fmt.Printf("Vehículo %d está saliendo del estacionamiento.\n", vehicle.ID)
	p.CurrentCount--
}
