package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type FyneAdapter struct {
	window fyne.Window
	status *widget.Label
}

func NewFyneAdapter() *FyneAdapter {
	a := app.New()
	w := a.NewWindow("Simulador de Estacionamiento")
	status := widget.NewLabel("Estado del estacionamiento")

	w.Resize(fyne.NewSize(800, 800))

	return &FyneAdapter{
		window: w,
		status: status,
	}
}

func (f *FyneAdapter) Show() {
	f.window.ShowAndRun()
}

func (f *FyneAdapter) UpdateStatus(status string) {
	if f.status == nil {
		return
	}

	if len(status) > 100 {
		status = status[:100] + "..."
	}

	f.status.SetText(status)
}

func (f *FyneAdapter) AddLog(text string) {
	println(text)
}
