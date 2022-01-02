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
	w := a.NewWindow("Circle")
	// Resize window
	w.Resize(fyne.NewSize(400, 400))
	//circle
	circle1 := canvas.NewCircle(color.Black)

	w.SetContent(circle1)
	w.ShowAndRun()
}
