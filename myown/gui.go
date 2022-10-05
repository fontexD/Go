package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MovieResults struct {
	Results []Movie `json:"results"`
}

type Movie struct {
	Name string `json:"Name"`
}

func LoadMovies() (MovieResults, error) {
	data, err := ioutil.ReadFile("./buttons.json")
	if err != nil {
		return MovieResults{}, err
	}
	var movieResults MovieResults
	err = json.Unmarshal(data, &movieResults)
	if err != nil {
		return MovieResults{}, err
	}
	return movieResults, nil
}

func main() {

	moviesResults, err := LoadMovies()
	if err != nil {
		panic(err)
	}
	fmt.Printf("movies: %s\n", moviesResults)

	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	W := a.NewWindow("Application-OutSight")
	W.Resize(fyne.NewSize(640, 460))

	listView := widget.NewList(func() int {
		return len(moviesResults.Results)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("template")
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		object.(*widget.Label).Text = moviesResults.Results[id].Name

	})

	contenttext := widget.NewLabel("Welcome to this App")
	listView.OnSelected = func(id widget.ListItemID) {
		contenttext.Text = moviesResults.Results[id].Name
	}

	split := (container.NewHSplit(
		listView,
		container.NewMax(contenttext),
	))
	split.Offset = 0.2
	W.SetContent(split)
	W.ShowAndRun()

}
