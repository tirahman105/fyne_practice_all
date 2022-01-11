package main

// importing fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	// creating new app
	a := app.New()
	// creating new window and title
	w := a.NewWindow("Menu")
	w.Resize(fyne.NewSize(400, 400))

	// first item
	menuItem := &fyne.Menu{
		Label: "File",
		Items: nil,
	}
	menu := fyne.NewMainMenu(menuItem)

	w.SetMainMenu(menu)
	// show and run
	w.ShowAndRun()
}
