package main

// importing fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// creating new app
	a := app.New()
	// creating new window and title
	w := a.NewWindow("Card")
	w.Resize(fyne.NewSize(400, 400))

	//card

	card := widget.NewCard(
		"Md. Tareq Ibna Rahman",
		"BDS, PGT, Leturer, PDC",
		canvas.NewImageFromFile("./card_img.png"),
	)

	// setup content
	w.SetContent(
		card,
	)
	// show and run
	w.ShowAndRun()
}
