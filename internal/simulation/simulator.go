package simulation

import (
	"parking-simulator/internal/domain"
	"math"
	"math/rand"
	"time"
)

type EventType string

const (
	EventWaiting EventType = "waiting"
	EventEnter   EventType = "enter"
	EventExit    EventType = "exit"
)

type SimulationEvent struct {
	VehicleID int
	EventType EventType
	SpaceID   int
}

type Simulator struct {
	VehicleCount int
	ParkingLot   *domain.ParkingLot
	EventChan    chan SimulationEvent
	StopChan     chan struct{}
	IsRunning    bool
}

func NewSimulator(parkingLot *domain.ParkingLot, vehicleCount int) *Simulator {
	return &Simulator{
		ParkingLot:   parkingLot,
		VehicleCount: vehicleCount,
		EventChan:    make(chan SimulationEvent, 100),
		StopChan:     make(chan struct{}),
	}
}

func (s *Simulator) Run() {
	if s.IsRunning {
		return
	}
	s.IsRunning = true

	vehicleID := 0
	for s.IsRunning && (s.VehicleCount == 0 || vehicleID < s.VehicleCount) {
		select {
		case <-s.StopChan:
			s.IsRunning = false
			return
		default:
			vehicle := domain.NewVehicle(vehicleID)
			go s.processVehicle(vehicle)
			s.poissonDelay()
			vehicleID++
		}
	}
}

func (s *Simulator) Stop() {
	if s.IsRunning {
		s.StopChan <- struct{}{}
		s.IsRunning = false
	}
}

func (s *Simulator) processVehicle(vehicle *domain.Vehicle) {
	// Intentar estacionar
	for !s.ParkingLot.EnterVehicle(vehicle) {
		s.EventChan <- SimulationEvent{
			VehicleID: vehicle.ID,
			EventType: EventWaiting,
		}
		time.Sleep(time.Second)
	}

	s.EventChan <- SimulationEvent{
		VehicleID: vehicle.ID,
		EventType: EventEnter,
		SpaceID:   vehicle.GetParkingSpace(),
	}

	vehicle.StayParked()

	spaceID := vehicle.GetParkingSpace()
	s.ParkingLot.ExitVehicle(vehicle)
	s.EventChan <- SimulationEvent{
		VehicleID: vehicle.ID,
		EventType: EventExit,
		SpaceID:   spaceID,
	}
}

func (s *Simulator) poissonDelay() {
	lambda := 8.0
	u := rand.Float64()
	interArrivalTime := -math.Log(1.0-u) / lambda
	time.Sleep(time.Duration(interArrivalTime * float64(time.Second)))
}
