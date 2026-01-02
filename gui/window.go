package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	windowName   string  = "L System"
	windowWidth  float32 = 1200
	windowHeight float32 = 1200
	minimumIteration float64 = 0
	maximumIteration float64 = 9
)

func BuildMainWindow(app fyne.App, binder *Binder) fyne.Window {
	window := app.NewWindow(windowName)
	window.Resize(fyne.NewSize(windowWidth, windowHeight))

	leftPane := buildLeftPaneForm(binder)
	rightPane := buildRightPane(binder)

	split := container.NewHSplit(leftPane, rightPane)
	split.Offset = 0.2

	window.SetContent(split)

	return window
}

func buildLeftPaneForm(binder *Binder) fyne.CanvasObject {
	iterationsSlider := widget.NewSlider(minimumIteration, maximumIteration)
	iterationsSlider.SetValue(minimumIteration)
	iterationsSlider.Step = 1

	iterationsSlider.OnChanged = binder.OnChangedIterations

	form := widget.NewForm(
		widget.NewFormItem("Iterations", iterationsSlider),
	)

	return container.NewVBox(form)
}

func buildRightPane(binder *Binder) fyne.CanvasObject {
	binder.canvasImage.FillMode = canvas.ImageFillContain

	return container.NewStack(binder.canvasImage)
}
