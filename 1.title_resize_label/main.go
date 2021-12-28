package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// this code will create a new app
	a := app.New()

	// now we will create a window

	w := a.NewWindow("Hello Tareq")

	//resize our app window

	w.Resize(fyne.NewSize(600, 500))

	//widget

	w.SetContent(widget.NewLabel("Hello Tareq"))

	w.ShowAndRun()
}
