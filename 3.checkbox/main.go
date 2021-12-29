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

	w.Resize(fyne.NewSize(400, 500))

	//checkbox widget

	checkbox1 := widget.NewCheck("Checkbox", func(b bool) {
		if b == true {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}

	})

	w.SetContent(checkbox1)
	w.ShowAndRun()
}
