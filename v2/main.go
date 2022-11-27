package main

import (
	"encoding/json"
	"fmt"
	"log"
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

// Func to pull data from  http.get, it takes Urls as input/call , Returns
func Getdata(Urls string) (Data, error) {

	var JsonOutput Data
	rsp, err := http.Get(Urls)
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()

	err = json.NewDecoder(rsp.Body).Decode(&JsonOutput)
	if err != nil {
		log.Fatal(err)
	}
	return JsonOutput, err
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
			fmt.Printf("hej")
			text = widget.NewLabel(dataContainer().Value)
			contain2 := container.NewMax(text)

			split := (container.NewHSplit(
				listView,
				contain2,
			))
			split.Offset = 0.2

			dataContainer()
			W.SetContent(split)
			W.Show()
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

func dataContainer() Data {
	var JsonOutput Data
	var err error

	Urls := []string{"http://127.0.0.1:8080/health", "http://127.0.1.1:8080/health", "http://127.0.2.1:8080/health"}
	for x := range Urls {
		JsonOutput, err = Getdata(Urls[x])
		//error checking
		if err != nil {
			fmt.Println(err)
		}
	}
	return JsonOutput
}
