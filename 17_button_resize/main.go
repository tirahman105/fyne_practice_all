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
	w := a.NewWindow("resize button")
	w.Resize(fyne.NewSize(400, 400))

	// 1st widget entry

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter Your Name")
	entry.Resize(fyne.NewSize(200, 50))
	entry.Move(fyne.NewPos(40, 50))

	entry_email := widget.NewEntry()
	entry_email.SetPlaceHolder("Enter Your email")
	entry_email.Resize(fyne.NewSize(200, 50))
	entry_email.Move(fyne.NewPos(40, 100))

	entry_address := widget.NewEntry()
	entry_address.SetPlaceHolder("Enter Your Address")
	entry_address.Resize(fyne.NewSize(200, 50))
	entry_address.Move(fyne.NewPos(40, 150))

	// Button

	btn_submit := widget.NewButton("Submit", nil)
	btn_submit.Resize(fyne.NewSize(100, 50))
	btn_submit.Move(fyne.NewPos(40, 200))

	w.SetContent(
		container.NewWithoutLayout(
			entry,
			entry_email,
			entry_address,
			btn_submit,
		),
	)
	// show and run
	w.ShowAndRun()
}
