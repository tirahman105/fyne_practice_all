package main

import (
	"fmt"

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

	w.Resize(fyne.NewSize(60, 50))

	//button add

	btn := widget.NewButton("Click here", func() {
		fmt.Println("Go button")
	})

	w.SetContent(btn)
	w.ShowAndRun()
}
