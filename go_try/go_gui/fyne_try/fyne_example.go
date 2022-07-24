package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"sync"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	go lockResource(hello)
	vBox := container.NewVBox(
		container.NewHBox(widget.NewCard("one", "two", widget.NewLabel("test for hbox"))),
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
		widget.NewButton("Hide", func() {
			w.Hide()
			time.Sleep(time.Second * 5)
			w.Show()
		}),
	)
	w.SetContent(vBox)
	w.ShowAndRun()
}

var rwLock = sync.RWMutex{}

func lockResource(val interface{}) {
	for {
		time.Sleep(5 * time.Second)
		rwLock.Lock()
		val.(*widget.Label).Text = time.Now().Format(time.RFC3339)
		fmt.Println("update time title:[" + val.(*widget.Label).Text)
		val.(*widget.Label).Refresh()
		rwLock.Unlock()
	}
}
