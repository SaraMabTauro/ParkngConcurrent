package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SpaceState struct {
	IsOccupied bool
	VehicleID  int
}

type Components struct {
	StartButton *widget.Button
	StopButton  *widget.Button
	Spaces      []*fyne.Container
	Container   *fyne.Container
	carImages   []*canvas.Image
	parkingImg  *canvas.Image
}

func NewComponents(capacity int) *Components {
	parkingImg := canvas.NewImageFromFile("assets/estacionamiento.png")
	parkingImg.Resize(fyne.NewSize(120, 120))

	carImages := []*canvas.Image{
		canvas.NewImageFromFile("assets/azul.png"),
		canvas.NewImageFromFile("assets/rojo.png"),
		canvas.NewImageFromFile("assets/verde.png"),
	}

	for _, img := range carImages {
		img.Resize(fyne.NewSize(100, 100))
	}

	spaces := make([]*fyne.Container, capacity)
	for i := range spaces {
		parkingSpace := canvas.NewImageFromFile("assets/estacionamiento.png")
		parkingSpace.Resize(fyne.NewSize(120, 120))
		spaces[i] = container.NewWithoutLayout(parkingSpace)
	}

	c := &Components{
		StartButton: widget.NewButton("Iniciar Simulación", nil),
		StopButton:  widget.NewButton("Detener Simulación", nil),
		Spaces:      spaces,
		carImages:   carImages,
		parkingImg:  parkingImg,
	}

	return c
}

func (c *Components) SetCallbacks(onStart, onStop func()) {
	c.StartButton.OnTapped = onStart
	c.StopButton.OnTapped = onStop
}

func (c *Components) EnableControls(enable bool) {
	c.StartButton.Enable()
	c.StopButton.Enable()
}
