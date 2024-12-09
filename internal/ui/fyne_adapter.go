package ui

import (
	"parking-simulator/pkg/simulation"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type FyneAdapter struct {
	window fyne.Window
	status *widget.Label
	logs   *widget.Label
	simulator *simulation.Simulator
	carImages []*canvas.Image
}

func NewFyneAdapter() *FyneAdapter {
	a := app.New()
	w := a.NewWindow("Simulador de Estacionamiento")

	status := widget.NewLabel("Estado del estacionamiento")
	logs := widget.NewLabel("Log de Eventos")

	carImages := make([]*canvas.Image, 20)
	for i := range carImages {
		carImages[i] = canvas.NewImageFromFile("../img/car.png")
		carImages[i].FillMode = canvas.ImageFillContain
		carImages[i].Hidden = true
	}

	canvasObjects := make([]fyne.CanvasObject, len(carImages))
	for i, img := range carImages {
		canvasObjects[i] = img
	}

	parkingSpaces := container.NewGridWithColumns(5, canvasObjects...)

	content := container.NewVBox(
		status,
		parkingSpaces,
		logs,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 300))

	return &FyneAdapter{
		window: w,
		status: status,
		logs:   logs,
		carImages: carImages,
	}
}

func (f *FyneAdapter) Show() {
	f.window.ShowAndRun()
}

func (f *FyneAdapter) UpdateStatus(text string) {
	f.status.SetText(text)
}

func (f *FyneAdapter) AddLog(text string) {
	currentText := f.logs.Text
	if len(currentText) > 1000 { // tamaño del log
		currentText = currentText[len(currentText)-1000:]
	}
	f.logs.SetText(currentText + "\n" + text)
}
