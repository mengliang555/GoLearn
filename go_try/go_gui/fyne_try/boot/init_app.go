package boot

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var appRoot fyne.App

func GetRootApp() fyne.App {
	return appRoot
}

func init() {
	appRoot = app.New()
}
