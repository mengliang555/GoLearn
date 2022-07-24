package step_one

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"testing"
)

func initFunc(object fyne.CanvasObject) {
	a := app.New()
	w := a.NewWindow("Try")
	w.SetContent(container.NewVBox(
		object,
	))
	w.ShowAndRun()
}

func TestGetAppTitle(t *testing.T) {
	initFunc(GetAppTitle())
}
