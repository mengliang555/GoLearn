package step_one

import (
	"fyne.io/fyne/v2/widget"
)

func PasswordTest() *widget.TextGrid {
	return widget.NewTextGrid()
}

func GetAButtonWithFuncAndString(handler func(), info string) *widget.Button {
	return widget.NewButton(info, handler)
}

