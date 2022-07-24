package step_one

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)


func GetAppTitle(item ...fyne.CanvasObject) *fyne.Container {
	if len(item) <= 0 {
		panic("make sure the title line has item")
	}
	return container.NewHBox(item...)
}

func RegisterCanvasObjectToLayout(layout fyne.Layout, item []fyne.CanvasObject) *fyne.Container {
	return container.New(layout, item...)
}
