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
	w := a.NewWindow("Child_menu")
	w.Resize(fyne.NewSize(400, 400))

	// menu item
	menuItem1 := fyne.NewMenuItem("edit", nil)
	menuItem2 := fyne.NewMenuItem("Details", nil)
	menuItem3 := fyne.NewMenuItem("Home", nil)

	//child menu

	menuItem1.ChildMenu = fyne.NewMenu(
		"",
		fyne.NewMenuItem("Copy", nil),
		fyne.NewMenuItem("Cut", nil),
		fyne.NewMenuItem("Paste", nil),
	)
	menuItem2.ChildMenu = fyne.NewMenu(
		"",
		fyne.NewMenuItem("Copy", nil),
		fyne.NewMenuItem("Cut", nil),
		fyne.NewMenuItem("Paste", nil),
	)

	//  New menu

	newMenu1 := fyne.NewMenu("File", menuItem1, menuItem2, menuItem3)
	newMenu2 := fyne.NewMenu("Help", menuItem1, menuItem2, menuItem3)

	menu := fyne.NewMainMenu(newMenu1, newMenu2)

	w.SetMainMenu(menu)
	// show and run
	w.ShowAndRun()
}
