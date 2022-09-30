//My First GO Program for monittoring web-site

package main

import (
	"encoding/json"
	"image/color"
	"log"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Data struct {
	Name  string `json:"key"`
	Value string `json:"value"`
}

var CheckInterval int = 5
var Url string = "http://127.0.0.1:8080/health"

func getData(url string) (Data, error) {
	var fact Data
	resp, err := http.Get(url)
	if err != nil {
		return Data{}, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&fact)
	if err != nil {
		return Data{}, err
	}
	return fact, nil
}

func main() {
	var fact string
	var url string
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Applicatiton Out-Sight")
	w.Resize(fyne.NewSize(800, 300))

	button0 := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		log.Println("tapped home")
	})
	image := canvas.NewImageFromFile("healthy.png")
	image.SetMinSize(fyne.Size{Width: 30, Height: 30})
	image.FillMode = canvas.ImageFill(canvas.ImageScalePixels)
	text1 := widget.NewLabel("")
	text0 := widget.NewLabel("Welcome To This App V1.0")
	button1 := widget.NewButtonWithIcon("Spa Status", theme.InfoIcon(), func() {
		url = "http://127.0.0.1:8000/health"
		doTheThing(url)
		text1 := canvas.NewText(url, color.White)
		text2 := canvas.NewText(fact, color.White)
		image1 := canvas.NewImageFromFile("unhealthy.png")
		image1.SetMinSize(fyne.Size{Width: 30, Height: 30})
		image1.FillMode = canvas.ImageFill(canvas.ImageScalePixels)
		grid := container.New(layout.NewGridLayout(3), text1, text2, image1)

		side := container.New(layout.NewVBoxLayout(), button0)
		content2 := container.New(layout.NewHBoxLayout(), side, widget.NewSeparator(), grid)
		w.SetContent(content2)
		w.Show()

	})
	side := container.New(layout.NewVBoxLayout(), button0, button1)
	content := container.New(layout.NewHBoxLayout(), side, widget.NewSeparator(), image, text0, text1)

	//split := container.NewHSplit(side, content)
	//split.Offset = 0.25

	w.SetContent(content)

	w.ShowAndRun()

}

var url string = ""

func doTheThing(string url) (Data, Error) {

	url = ""
	var checkInterval int = 5
	for range time.Tick(checkInterval * time.Second) {
		fact, err := getData(url)
		if err != nil {
			log.Fatal(err)
		}

	}
}
