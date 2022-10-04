package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	Loadui()
}

func Loadui() {
	a := app.New()

	w := a.NewWindow("Fyne Demo")

	w.SetMaster()

	container()
	split := container.NewHSplit(makeNav(), tutorial)

	split.Offset = 0
	w.SetContent(split)

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}

func container() {

	content := container.NewMax()
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	intro.Wrapping = fyne.TextWrapWord
	tutorial := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)
	return
}
func makeNav() fyne.CanvasObject {

	tree := widget.NewTreeWithStrings(menuItems)

	return container.NewBorder(nil, nil, nil, nil, tree)
}

var menuItems = map[string][]string{
	"":            {"welcome", "collections", "advanced"},
	"collections": {"list", "table"},
}
