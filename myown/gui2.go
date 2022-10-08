package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
<<<<<<< HEAD
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func makeAlphabet() []string {
	var alphabet []string
	for ch := 'A'; ch <= 'Z'; ch++ {
		alphabet = append(alphabet, string(ch))
	}
	return alphabet
}

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
	app := app.New()
	window := app.NewWindow("tac_hub")
	window.Resize(fyne.NewSize(200, 200))

	rawData := []string{"Home", "app"}
	data := binding.BindStringList(&rawData)
	list := widget.NewListWithData(
		data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		fmt.Printf(rawData[1])
	}

	window.SetContent(list)
	window.ShowAndRun()
=======
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
>>>>>>> 844b52b711ea828d6a7acfae7cf94b9efa26bbdb
}
