package main

// importing fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// creating new app
	a := app.New()
	// creating new window and title
	w := a.NewWindow("Layout NewVBox and NewHbox")
	w.Resize(fyne.NewSize(400, 400))
	// New button
	btn1 := widget.NewButton("click me", func() {
	})
	label1 := widget.NewLabel("here is my text")
	// NewHBox
	// box1 := container.NewHBox( // Horizontal
	// 	btn1,
	// 	label1,
	// )
	// NewVBox
	box1 := container.NewVBox( // Vertical
		label1,
		btn1,
	)
	// setup content
	w.SetContent(
		box1,
	)
	// show and run
	w.ShowAndRun()
}
