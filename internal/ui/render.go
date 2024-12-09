package ui

import (
	"fmt"
	"parking-simulator/pkg/simulation"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Renderer struct {
	Adapter    *FyneAdapter
	Simulator  *simulation.Simulator
	Components *Components
	stopChan   chan struct{}
}

func NewRenderer(adapter *FyneAdapter, simulator *simulation.Simulator) *Renderer {
	r := &Renderer{
		Adapter:    adapter,
		Simulator:  simulator,
		Components: NewComponents(nil, nil),
		stopChan:   make(chan struct{}),
	}

	r.Components = NewComponents(r.StartSimulation, r.StopSimulation)

	canvasObjects := make([]fyne.CanvasObject, len(adapter.carImages))
	for i, img := range adapter.carImages {
		canvasObjects[i] = img
	}

	adapter.window.SetContent(
		container.NewVBox(
			adapter.status,
			container.NewGridWithColumns(5, canvasObjects...),
			r.Components.StartButton,
			r.Components.StopButton,
		),
	)

	return r
}

func (r *Renderer) StartSimulation() {
	go func() {
		for i := 0; i < r.Simulator.VehicleCount; i++ {
			select {
			case <-r.stopChan:
				return
			default:
				vehicle := simulation.NewVehicle(i)
				go r.handleVehicle(vehicle)
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
}

func (r *Renderer) StopSimulation() {
	r.stopChan <- struct{}{}
	r.Adapter.AddLog("Simulación detenida.")
}

func (r *Renderer) handleVehicle(vehicle *simulation.Vehicle) {
	spaceIndex := vehicle.ID % 20

	for !r.Simulator.ParkingLot.EnterVehicle(vehicle) {
		r.Adapter.AddLog(fmt.Sprintf("Vehículo %d esperando por espacio...", vehicle.ID))
		time.Sleep(time.Second)
	}

	r.Adapter.carImages[spaceIndex].Hidden = false
	r.Adapter.UpdateStatus(fmt.Sprintf("Vehículo %d entró al estacionamiento", vehicle.ID))
	r.Adapter.AddLog(fmt.Sprintf("Vehículo %d estacionado.", vehicle.ID))

	vehicle.StayParked()

	time.Sleep(time.Second * 5)

	r.Simulator.ParkingLot.ExitVehicle(vehicle)
	r.Adapter.carImages[spaceIndex].Hidden = true
	r.Adapter.UpdateStatus(fmt.Sprintf("Vehículo %d salió del estacionamiento", vehicle.ID))
	r.Adapter.AddLog(fmt.Sprintf("Vehículo %d ha salido del estacionamiento.", vehicle.ID))
}
