package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TapLabel struct {
	*widget.Label //composition

	//function pointers to set to get events
	OnTapped func(string)
}

func (mc *TapLabel) Tapped(pe *fyne.PointEvent) {
	if mc.OnTapped != nil {
		mc.OnTapped(mc.Text)
	}
}

func NewTapLabel(text string, tappedLeft func(string)) *TapLabel {
	return &TapLabel{widget.NewLabel(text), tappedLeft}
}

func alphabetToBrands(letter string) {
	fmt.Println(letter)
}

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
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*TapLabel).SetText(data[i])
		})
}
