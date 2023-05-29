package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"go_gui/server"
	"go_gui/tool"
)

func login(w fyne.Window) {
	if tool.Api.Token.Get() != "" {
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
				newHome(w)
				return
			}
			di := dialog.NewInformation("警告！", "賬號密碼錯誤！", w)
			di.Show()
		},
	}

	w.SetContent(container.NewVBox(form))
}
