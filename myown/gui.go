package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Menubar struct {
	Name string
}

func main() {

	Menu := []string{"Home", "App-Status", "exit"}

	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	W := a.NewWindow("Application-OutSight")
	W.Resize(fyne.NewSize(640, 460))
	W.SetMaster()

	listView := widget.NewList(func() int {
		return len(Menu)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("template")
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		object.(*widget.Label).Text = Menu[id]
	})

	W.SetContent(container.NewHSplit(
		listView,
		container.NewMax(),
	))

	W.ShowAndRun()

}
