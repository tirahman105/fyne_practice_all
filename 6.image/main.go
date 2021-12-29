package main

// import Fyne
import (
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
	//img widget

	img := canvas.NewImageFromFile("./golang.png")

	w.SetContent(img)
	w.ShowAndRun()
}
