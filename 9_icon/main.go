package main

// importing Fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// creating app
	a := app.New()
	// create window
	w := a.NewWindow("Icon")
	//resizing windows
	w.Resize(fyne.NewSize(200, 200))
	// Creating Icon Widget
	IconX := widget.NewIcon(theme.CancelIcon())
	// setup content
	w.SetContent(IconX)
	// show windows
	w.ShowAndRun()
}
