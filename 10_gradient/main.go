package main

// import Fyne
import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	// New App
	a := app.New()
	// New Window and Title
	w := a.NewWindow("Gradient : Canvas")
	w.Resize(fyne.NewSize(400, 400))
	// Radial Gradient
	// gradient1 := canvas.NewRadialGradient(
	// 	color.White, color.NRGBA{R: 255, G: 0, B: 0, A: 255}) // 45 degree angle
	// Linear Gradient
	gradient2 := canvas.NewLinearGradient(
		color.White, color.NRGBA{R: 255, G: 0, B: 0, A: 255}, 270) // 45 degree angle
	// Horizontal Gradient
	//gradient3 := canvas.NewHorizontalGradient(
	// color.White, color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	// Vertical Gradient
	//gradient4 := canvas.NewVerticalGradient(
	// color.White, color.NRGBA{R: 255, G: 0, B: 0, A: 255}) // 45 degree angle
	// setup content
	w.SetContent(gradient2)
	w.ShowAndRun()
}
