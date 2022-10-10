package main

import (
	"encoding/json"
	"fmt"
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

var Menu = []string{"Home", "App-Status"}

// Define strucutre of data
type Data struct {
	Name  string `json:"key"`
	Value string `json:"value"`
}

func main() {
	var W fyne.Window
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	W = a.NewWindow("Application-OutSight")
	W.Resize(fyne.NewSize(640, 480))
	text := widget.NewLabel("W½elcome to This App")

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
			text := widget.NewLabel("Storage.plan2learn.dk")
			text2 := widget.NewLabel(dataReturn().Value)
			text3 := canvas.NewImageFromFile("./healthy.png")
			text3.FillMode = canvas.ImageFillContain
			text3.SetMinSize(fyne.Size{Width: 30, Height: 20})
			grid := container.New(layout.NewGridWrapLayout(fyne.Size{Width: 200, Height: 50}), text, text2, text3)

			split := (container.NewHSplit(
				listView,
				grid,
			))
			split.Offset = 0.2
			W.SetContent(split)
			go func() {
				for range time.Tick(time.Second * 5) {
					text2.SetText(dataReturn().Value)

				}
			}()
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

func dataReturn() Data {
	//Define var Jsonoutput to what Datastructure it holds
	var JsonOutput Data
	var err error

	// Create array of urls to send t function GetData
	Urls := []string{"http://192.168.10.11:8080/health"}

	// Infinity loop which render every 5 second
	//loop:
	for x := range Urls {
		JsonOutput, err = Getdata(Urls[x])
		//error checking
		if err != nil {
			log.Fatal(err)
		}
		// print .Value from Jsonoutput Struct variable
	}
	return JsonOutput
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
