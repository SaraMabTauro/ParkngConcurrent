package ui

import (
	"parking-simulator/internal/simulation"
	"fmt"
	"math/rand"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type RenderState struct {
	IsRunning bool
	mutex     sync.Mutex
}

type Renderer struct {
	Adapter    *FyneAdapter
	Simulator  *simulation.Simulator
	Components *Components
	state      *RenderState
	stopChan   chan struct{}
}

func NewRenderer(adapter *FyneAdapter, simulator *simulation.Simulator) *Renderer {
	r := &Renderer{
		Adapter:    adapter,
		Simulator:  simulator,
		Components: NewComponents(simulator.ParkingLot.Capacity),
		stopChan:   make(chan struct{}),
		state:      &RenderState{},
	}

	r.setupUI()
	return r
}

func (r *Renderer) setupUI() {
	r.Components.StartButton.OnTapped = r.StartSimulation
	r.Components.StopButton.OnTapped = r.StopSimulation

	r.Adapter.window.Resize(fyne.NewSize(800, 600))

	mainGrid := container.NewGridWithRows(4)
	for row := 0; row < 4; row++ {
		rowContainer := container.NewGridWithColumns(5)
		for col := 0; col < 5; col++ {
			index := row*5 + col
			if index < len(r.Components.Spaces) {
				spacePadded := container.NewPadded(r.Components.Spaces[index])
				rowContainer.Add(spacePadded)
			}
		}
		mainGrid.Add(rowContainer)
	}

	buttonContainer := container.NewCenter(
		container.NewHBox(
			r.Components.StartButton,
			widget.NewSeparator(),
			r.Components.StopButton,
		),
	)

	content := container.NewBorder(
		r.Adapter.status,
		buttonContainer,
		nil,
		nil,
		mainGrid,
	)

	r.Adapter.window.SetContent(content)
}

func (r *Renderer) StartSimulation() {
	r.state.mutex.Lock()
	if r.state.IsRunning {
		r.state.mutex.Unlock()
		return
	}
	r.state.IsRunning = true
	r.state.mutex.Unlock()

	go r.Simulator.Run()
	go r.handleSimulationEvents()
}

func (r *Renderer) handleSimulationEvents() {
	for event := range r.Simulator.EventChan {
		switch event.EventType {
		case "waiting":
			r.Adapter.AddLog(fmt.Sprintf("Vehículo %d esperando...", event.VehicleID))
		case "enter":
			if event.SpaceID >= 0 && event.SpaceID < len(r.Components.Spaces) {
				carIndex := rand.Intn(3)
				carImg := canvas.NewImageFromFile(fmt.Sprintf("assets/%s.png",
					[]string{"azul", "rojo", "verde"}[carIndex]))
				carImg.Resize(fyne.NewSize(90, 90))
				carImg.Move(fyne.NewPos(15, 15))

				r.Components.Spaces[event.SpaceID].Add(carImg)
				r.Components.Spaces[event.SpaceID].Refresh()
			}
		case "exit":
			if event.SpaceID >= 0 && event.SpaceID < len(r.Components.Spaces) {
				space := r.Components.Spaces[event.SpaceID]
				objects := space.Objects
				if len(objects) > 1 {
					space.Remove(objects[len(objects)-1])
					space.Refresh()
				}
			}
		}
	}
}

func (r *Renderer) StopSimulation() {
	r.state.mutex.Lock()
	if !r.state.IsRunning {
		r.state.mutex.Unlock()
		return
	}
	r.state.IsRunning = false
	r.state.mutex.Unlock()

	r.Simulator.Stop()
	r.Adapter.AddLog("Simulación detenida.")
}
