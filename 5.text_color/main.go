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
	w := a.NewWindow("Tareq")
	// Resize window
	w.Resize(fyne.NewSize(400, 400))
	// Our Text widget in Canvas
	//first value for label/cation
	// 2nd for color value
	// Now creating Color
	color := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	// R is Red and Range is zero to 255
	// B is blue and G is Green
	// A is for alpha for opacity
	text := canvas.NewText("Dr. Tareq", color)
	// Setup widget.
	w.SetContent(text)
	w.ShowAndRun()
}
