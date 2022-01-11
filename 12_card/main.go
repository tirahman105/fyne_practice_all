package main

// importing fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// creating new app
	a := app.New()
	// creating new window and title
	w := a.NewWindow("Card")
	w.Resize(fyne.NewSize(400, 400))

	//card

	card1 := widget.NewCard(
		"Md. Tareq Ibna Rahman",
		"BDS, PGT, Leturer, PDC",
		canvas.NewImageFromFile("./card_img.png"),
	)
	btn1 := widget.NewButton("Click", func() {

	})
	label1 := widget.NewLabel("This is test")
	label2 := widget.NewLabel("This is test2")

	// setup content
	w.SetContent(
		container.NewHBox(
			card1,
			container.NewVBox(
				label1,
				label2,
				btn1,
			),
		),
	)
	// show and run
	w.ShowAndRun()
}
