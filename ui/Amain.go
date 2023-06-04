package ui

import (
	"fyne.io/fyne/v2"
	app2 "fyne.io/fyne/v2/app"
)

func Main() {
	app := app2.New()
	wApp := app.NewWindow("TChat")
	login(wApp)
	wApp.Resize(fyne.NewSize(840, 460))
	wApp.ShowAndRun()

}
