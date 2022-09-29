package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var content fyne.CanvasObject
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Fyne Demo")

	w.SetMaster()

	containers := container.NewMax()
	title := widget.NewLabel("Application Status")
	intro := widget.NewLabel("Storage.plan2learn.dk")
	intro2 := widget.NewLabel("")
	intro2.SetText("hej")
	intro.Wrapping = fyne.TextWrapWord
	TABCONTENT := container.NewBorder(container.NewVBox(title, intro, intro2), nil, nil, containers)

	tabs := container.NewAppTabs(
		container.NewTabItem("TABNAME", TABCONTENT),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	content = container.NewBorder(
		container.NewVBox(title, intro, intro2), nil, nil, containers)

	split := container.NewHSplit(tabs, content)

	split.Offset = 0
	w.SetContent(split)

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}
