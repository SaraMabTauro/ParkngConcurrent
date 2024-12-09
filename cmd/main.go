package main

import (
	"parking-simulator/internal/ui"
	"parking-simulator/pkg/simulation"
)

func main() {
	capacity := 20
	vehicleCount := 100

	simulator := simulation.NewSimulator(capacity)
	simulator.VehicleCount = vehicleCount

	adapter := ui.NewFyneAdapter()
	renderer := ui.NewRenderer(adapter, simulator)
	renderer.Adapter.Show()
}
