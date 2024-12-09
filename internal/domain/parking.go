package domain

import (
	"parking-simulator/pkg/concurrencia"
	"fmt"
	"sync"
	"time"
)

const (
	DirectionEnter = "entrada"
	DirectionExit  = "salida"
	DirectionNone  = "ninguno"
)

type ParkingLot struct {
	Capacity           int
	CurrentCount       int
	EntrySemaphore     *semaphore.Semaphore
	Spaces             []bool
	Direction          string
	WaitingCount       int
	directionMutex     sync.Mutex
	sameDirectionCount int
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		Capacity:       capacity,
		Spaces:         make([]bool, capacity),
		EntrySemaphore: semaphore.NewSemaphore(1),
		Direction:      DirectionNone,
	}
}

func (p *ParkingLot) EnterVehicle(vehicle *Vehicle) bool {
	if p.CurrentCount >= p.Capacity {
		fmt.Printf("Estacionamiento lleno, el vehículo %d está esperando.\n", vehicle.ID)
		return false
	}

	p.directionMutex.Lock()
	currentDirection := p.Direction
	if currentDirection == DirectionExit {
		p.WaitingCount++
		p.directionMutex.Unlock()
		return false
	}

	if currentDirection == DirectionNone {
		p.Direction = DirectionEnter
		p.sameDirectionCount = 1
	} else {
		p.sameDirectionCount++
	}
	p.directionMutex.Unlock()

	if p.sameDirectionCount > 1 {
		spaceID := p.findFreeSpace()
		if spaceID != -1 {
			p.Spaces[spaceID] = true
			p.CurrentCount++
			vehicle.AssignParkingSpace(spaceID)
			fmt.Printf("Vehículo %d está entrando al espacio %d (mismo sentido).\n", vehicle.ID, spaceID)
			return true
		}
		return false
	}

	p.EntrySemaphore.Acquire()
	spaceID := p.findFreeSpace()
	if spaceID != -1 {
		p.Spaces[spaceID] = true
		p.CurrentCount++
		vehicle.AssignParkingSpace(spaceID)
		fmt.Printf("Vehículo %d está entrando al espacio %d.\n", vehicle.ID, spaceID)
		p.EntrySemaphore.Release()
		return true
	}

	p.EntrySemaphore.Release()
	return false
}

func (p *ParkingLot) ExitVehicle(vehicle *Vehicle) {
	p.directionMutex.Lock()
	if p.Direction == DirectionEnter && p.WaitingCount > 0 {
		p.directionMutex.Unlock()
		time.Sleep(time.Second)
		return
	}

	currentDirection := p.Direction
	if currentDirection == DirectionNone {
		p.Direction = DirectionExit
		p.sameDirectionCount = 1
	} else if currentDirection == DirectionExit {
		p.sameDirectionCount++
	}
	p.directionMutex.Unlock()

	p.EntrySemaphore.Acquire()
	spaceID := vehicle.GetParkingSpace()
	if spaceID >= 0 && spaceID < len(p.Spaces) {
		p.Spaces[spaceID] = false
		p.CurrentCount--
		fmt.Printf("Vehículo %d está saliendo del espacio %d.\n", vehicle.ID, spaceID)
		vehicle.RemoveParkingSpace()
	}
	p.EntrySemaphore.Release()

	p.directionMutex.Lock()
	p.sameDirectionCount--
	if p.sameDirectionCount == 0 {
		p.Direction = DirectionNone
		if p.WaitingCount > 0 {
			p.WaitingCount--
		}
	}
	p.directionMutex.Unlock()
}

func (p *ParkingLot) findFreeSpace() int {
	for i, occupied := range p.Spaces {
		if !occupied {
			return i
		}
	}
	return -1
}

func (p *ParkingLot) IsFull() bool {
	return p.CurrentCount >= p.Capacity
}
