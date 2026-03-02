package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	
	w.SetContent(widget.NewButton("Hi!", func() {
			hello.SetText("Welcome shiva :)")
		})
	)
	w.ShowAndRun()
}
