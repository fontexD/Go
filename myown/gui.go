package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var Menu = []string{"Home", "App-Status"}

func main() {
	contenttext := widget.NewLabel("Welcome to this App")
	loadUI(contenttext)
}

func loadUI(contenttext *widget.Label) {

	var W fyne.Window
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	W = a.NewWindow("Application-OutSight")
	W.Resize(fyne.NewSize(640, 460))

	split := (container.NewHSplit(
		menuBar(Menu),
		container.NewMax(contenttext),
	))
	split.Offset = 0.2
	W.SetContent(split)
	W.ShowAndRun()
}

func menuBar(Menu []string) *widget.List {
	listView := widget.NewList(func() int {
		return len(Menu)
	},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(Menu[id])
		})
	listView.OnSelected = func(id widget.ListItemID) {
		if id == 0 {
			contenttext := widget.NewLabel("newtext")
			loadUI(contenttext)
		} else if id == 1 {
			fmt.Println("app")
		}
		if id == 2 {
			fmt.Println("exit")
		}
	}
	return listView
}



func changeText(s *widget.Label) {

	var contenttext widget.Label
	contenttext = *widget.NewLabel("New-Value_pointer")
	*s = contenttext
}
