package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Закрытие окна")
	label := widget.NewLabel("Нажмите любую кнопку")
	button1 := widget.NewButton("Закрыть окно", func() {
		w.Close()
	})
	button2 := widget.NewButton("Нажми меня", func() {
		label.SetText("жмякнул")
	})
	buttons := container.NewHBox(button1, button2)
	w.SetContent(container.NewVBox(
		label,
		buttons,
	))

	w.SetOnClosed(func() {
		fmt.Println("Окно закрыто")
	})

	w.ShowAndRun()

}
