package main

// importing fyne
import (
	"fmt"

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
	menuItem1 := fyne.NewMenuItem("New", func() { fmt.Println("New pressed") })
	menuItem2 := fyne.NewMenuItem("Save", func() { fmt.Println("Save pressed") })
	menuItem3 := fyne.NewMenuItem("Edit", nil)

	//  New menu

	newMenu := fyne.NewMenu("File", menuItem1, menuItem2, menuItem3)

	menu := fyne.NewMainMenu(newMenu)

	w.SetMainMenu(menu)
	// show and run
	w.ShowAndRun()
}
