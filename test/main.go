package test

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("miniChat")
	var data []string
	for i := 0; i < 99; i++ {
		data = append(data, strconv.Itoa(i))
	}
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	intro.Wrapping = fyne.TextWrapWord
	content := container.NewMax()
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])

		})

	tutorial := container.NewBorder(container.NewVBox(widget.NewLabel("Component name"), widget.NewSeparator(), intro), nil, nil, nil, content)
	split := container.NewHSplit(list, tutorial)
	split.Offset = 0.3
	w.SetContent(split)
	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}
