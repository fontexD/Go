package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	W := a.NewWindow("Application-OutSight")
	W.Resize(fyne.NewSize(640, 460))

	var menu = []string{"Home", "App-Status", "Exit"}

	contenttext := widget.NewLabel("Welcome to this App")

	split := (container.NewHSplit(
		menuBar(menu),
		container.NewMax(contenttext),
	))
	split.Offset = 0.2
	W.SetContent(split)
	W.ShowAndRun()

}

func menuBar(data []string) *widget.List {
	return widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewPadded(
				widget.NewLabel("Will be replaced"),
				widget.NewButton("", nil),
			)
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(data[id])

			// new part
			o.(*fyne.Container).Objects[1].(*widget.Button).OnTapped.(data[id]) = func() {
				fmt.Println("I am button " + data[id])
			}
		},
	)
}
