package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
<<<<<<< HEAD
=======
	"time"
>>>>>>> a6f4352e77f32e2e95c103afd6856374f7a801be

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var Menu = []string{"Home", "App-Status"}

<<<<<<< HEAD
=======
// Define strucutre of data
>>>>>>> a6f4352e77f32e2e95c103afd6856374f7a801be
type Data struct {
	Name  string `json:"key"`
	Value string `json:"value"`
}

<<<<<<< HEAD
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

=======
>>>>>>> a6f4352e77f32e2e95c103afd6856374f7a801be
func main() {
	var W fyne.Window
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	W = a.NewWindow("Application-OutSight")
<<<<<<< HEAD
	W.Resize(fyne.NewSize(640, 460))
	text := widget.NewLabel("Welcome to This App")
=======
	W.Resize(fyne.NewSize(640, 480))
	text := widget.NewLabel("W½elcome to This App")
>>>>>>> a6f4352e77f32e2e95c103afd6856374f7a801be

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
<<<<<<< HEAD
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

=======

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
>>>>>>> a6f4352e77f32e2e95c103afd6856374f7a801be
	contain := container.NewMax(text)

	split := (container.NewHSplit(
		listView,
		contain,
	))
	split.Offset = 0.2
	W.SetContent(split)
	W.ShowAndRun()
}

<<<<<<< HEAD
func dataContainer() Data {
	var JsonOutput Data
	var err error

	Urls := []string{"http://127.0.0.1:8080/health", "http://127.0.1.1:8080/health", "http://127.0.2.1:8080/health"}
=======
func dataReturn() Data {
	//Define var Jsonoutput to what Datastructure it holds
	var JsonOutput Data
	var err error

	// Create array of urls to send t function GetData
	Urls := []string{"http://192.168.10.11:8080/health"}

	// Infinity loop which render every 5 second
	//loop:
>>>>>>> a6f4352e77f32e2e95c103afd6856374f7a801be
	for x := range Urls {
		JsonOutput, err = Getdata(Urls[x])
		//error checking
		if err != nil {
<<<<<<< HEAD
			fmt.Println(err)
		}
	}
	return JsonOutput
=======
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
>>>>>>> a6f4352e77f32e2e95c103afd6856374f7a801be
}
