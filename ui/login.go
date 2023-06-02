package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"go_gui/server"
)

func login(w fyne.Window) {
	if server.Api.LoginVerify.LoginConfig() {
		if !server.Api.Login.StartTcp() {

		}
		newHome(w)
		return
	}
	entry := widget.NewEntry()
	password := widget.NewPasswordEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "account", Widget: entry},
			{Text: "password", Widget: password},
		},
		OnSubmit: func() {
			if server.Api.Login.Verify(entry.Text, password.Text) {
				server.Api.Login.StartTcp()
				newHome(w)
				return
			}
			di := dialog.NewInformation("警告！", "賬號密碼錯誤！", w)
			di.Show()
		},
	}

	w.SetContent(container.NewVBox(form))
}
