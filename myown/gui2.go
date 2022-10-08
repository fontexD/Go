package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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
}
