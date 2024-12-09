package main

import (
	"parking-simulator/internal/domain"
	"parking-simulator/internal/ui"
	"parking-simulator/internal/simulation"
)

func main() {
	capacity := 20
	vehicleCount := 100

	parkingLot := domain.NewParkingLot(capacity)
	simulator := simulation.NewSimulator(parkingLot, vehicleCount)

	adapter := ui.NewFyneAdapter()
	renderer := ui.NewRenderer(adapter, simulator)

	renderer.Adapter.Show()
}