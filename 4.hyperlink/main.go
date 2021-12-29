package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// this code will create a new app
	a := app.New()

	// now we will create a window

	w := a.NewWindow("Hello Tareq")

	//resize our app window

	w.Resize(fyne.NewSize(400, 500))

	//add url

	url, _ := url.Parse("https://github.com/tirahman105")

	//hyperlink widget

	hyperlink := widget.NewHyperlink("Visit My Github", url)

	w.SetContent(hyperlink)
	w.ShowAndRun()
}
