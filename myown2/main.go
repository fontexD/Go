package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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

var Menu = []string{"Home", "App-Status"}

func updateTime(clock *widget.Label, s string) {

	formatted := (s)
	clock.SetText(formatted)

}
func main() {
	var W fyne.Window
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	W = a.NewWindow("Application-OutSight")
	W.Resize(fyne.NewSize(640, 460))
	//content := widget.NewLabel("Welcome to This App")
	s := "Welcome to This Appssss"
	clock := widget.NewLabel("Welcome to This App")
	contentcontainer := container.NewMax(clock)
	split := (container.NewHSplit(
		menuBar(Menu),
		contentcontainer,
	))
	split.Offset = 0.2
	W.SetContent(split)
	go func() {
		for range time.Tick(time.Second * 5) {
			updateTime(clock, s)
		}
	}()
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
			clock := widget.NewLabel("ssss")
			s := "Welcome to This App AGAIN!"
			updateTime(clock, s)
		} else if id == 1 {
			fmt.Println("app")
			clock := widget.NewLabel("ssss")
			s := "Welcome to This App AGAIN!"
			updateTime(clock, s)
		}
		if id == 2 {
			fmt.Println("exit")
		}
	}
	return listView
}
