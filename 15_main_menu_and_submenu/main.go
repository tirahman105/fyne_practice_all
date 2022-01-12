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

	// menu item
	menuItem1 := fyne.NewMenuItem("New", nil)
	menuItem2 := fyne.NewMenuItem("edit", nil)
	menuItem3 := fyne.NewMenuItem("run", nil)

	//  New menu

	newMenu1 := fyne.NewMenu("File", menuItem1, menuItem2, menuItem3)
	newMenu2 := fyne.NewMenu("Other", menuItem1, menuItem2, menuItem3)
	newMenu3 := fyne.NewMenu("Help", menuItem1, menuItem2, menuItem3)

	menu := fyne.NewMainMenu(newMenu1, newMenu2, newMenu3)

	w.SetMainMenu(menu)
	// show and run
	w.ShowAndRun()
}
