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
	w := a.NewWindow("Hsplit widget")
	w.Resize(fyne.NewSize(400, 400))

	// 1st widget
	label1 := widget.NewLabel("This is test")
	label2 := widget.NewLabel("This is test")

	//2nd widget
	btn1 := widget.NewButton("Click", func() {

	})

	//card

	// card1 := widget.NewCard(
	// 	"Md. Tareq Ibna Rahman",
	// 	"BDS, PGT, Leturer, PDC",
	// 	canvas.NewImageFromFile("./card_img.png"),
	// )
	// btn1 := widget.NewButton("Click", func() {

	// })
	// label1 := widget.NewLabel("This is test")
	// label2 := widget.NewLabel("This is test2")

	// // setup content
	// w.SetContent(
	// 	container.NewHBox(
	// 		card1,
	// 		container.NewVBox(
	// 			label1,
	// 			label2,
	// 			btn1,
	// 		),
	// 	),
	// )

	w.SetContent(
		container.NewHSplit(
			label1,
			container.NewVBox(
				label2,
				btn1,
			),
		),
	)
	// show and run
	w.ShowAndRun()
}
