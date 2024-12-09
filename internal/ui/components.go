package ui

import (

	"fyne.io/fyne/v2/widget"
)

type Components struct {
	StartButton *widget.Button
	StopButton  *widget.Button
}

func NewComponents(onStart func(), onStop func()) *Components {

	return &Components{
		StartButton: widget.NewButton("Iniciar Simulación", onStart),
		StopButton:  widget.NewButton("Detener Simulación", onStop),
	}
}
