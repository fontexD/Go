package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var Menu = []string{"Home", "App-Status"}

type Data struct {
	Name  string `json:"key"`
	Value string `json:"value"`
}

func getData() (Data, error) {
	var CheckInterval int = 5
	var Url string = "http://127.0.0.1:8080/health"
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
	var W fyne.Window
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	W = a.NewWindow("Application-OutSight")
	W.Resize(fyne.NewSize(640, 460))
	text := widget.NewLabel("Welcome to This App")

	// start container with welcome text
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
			text = widget.NewLabel("test2")
			contain2 := container.NewMax(text)

			split := (container.NewHSplit(
				listView,
				contain2,
			))
			split.Offset = 0.2
			W.SetContent(split)
		} else if id == 1 {
			fmt.Println("app")
		}
		if id == 2 {
			fmt.Println("exit")
		}
	}
	contain := container.NewMax(text)

	split := (container.NewHSplit(
		listView,
		contain,
	))
	split.Offset = 0.2
	W.SetContent(split)
	W.ShowAndRun()
}
