package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Application Out-Sight")

	w.SetMaster()

	content := container.NewMax()
	title := widget.NewLabel("Apllcation Status")
	intro := widget.NewLabel("https://Storage.Plan2learn.dk    Healthy")
	intro.Wrapping = fyne.TextWrapWord

	tutorial := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)

	split := container.NewHSplit(makeNav(), tutorial)

	split.Offset = 0
	w.SetContent(split)

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}

func makeNav() fyne.CanvasObject {

	a := fyne.CurrentApp()

	tree := widget.NewTreeWithStrings(menuItems)
	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree, layout.NewSpacer())
}

var menuItems = map[string][]string{
	"":            {"welcome", "collections", "advanced"},
	"collections": {"list", "table"},
}
