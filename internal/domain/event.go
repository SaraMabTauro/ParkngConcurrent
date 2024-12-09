package domain

import (
	"fmt"
)

type EventType string

const (
	EventWaiting EventType = "waiting"
	EventEnter   EventType = "enter"
	EventExit    EventType = "exit"
)

type ParkingEvent struct {
	VehicleID int
	EventType EventType
	SpaceID   int
}

func NewParkingEvent(vehicleID int, eventType EventType, spaceID int) *ParkingEvent {
	return &ParkingEvent{
		VehicleID: vehicleID,
		EventType: eventType,
		SpaceID:   spaceID,
	}
}

func (e *ParkingEvent) String() string {
	switch e.EventType {
	case EventWaiting:
		return fmt.Sprintf("Vehículo %d esperando por espacio", e.VehicleID)
	case EventEnter:
		return fmt.Sprintf("Vehículo %d entrando al espacio %d", e.VehicleID, e.SpaceID)
	case EventExit:
		return fmt.Sprintf("Vehículo %d saliendo del espacio %d", e.VehicleID, e.SpaceID)
	default:
		return fmt.Sprintf("Evento desconocido para vehículo %d", e.VehicleID)
	}
}
