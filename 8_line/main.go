package main

// import Fyne
import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	// Create new app
	a := app.New()
	// Create New Window and Title
	w := a.NewWindow("Line")
	// Resize window
	w.Resize(fyne.NewSize(400, 400))
	//creating line
	lineX := canvas.NewLine(color.White)

	w.SetContent(lineX)
	w.ShowAndRun()
}
