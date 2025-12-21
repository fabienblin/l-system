package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func buildMainWindow(app fyne.App, lsystem *Lsystem) fyne.Window {
	window := app.NewWindow(windowName)
	window.Resize(fyne.NewSize(windowWidth, windowHeight))

	leftPane := buildLeftPaneForm(lsystem)
	rightPane := buildRightPane(lsystem)

	split := container.NewHSplit(leftPane, rightPane)
	split.Offset = 0.2

	window.SetContent(split)

	return window
}

func buildLeftPaneForm(lsystem *Lsystem) fyne.CanvasObject {
	iterationsSlider := widget.NewSlider(1, 10)
	iterationsSlider.SetValue(1)
	iterationsSlider.Step = 1

	iterationsSlider.OnChanged = lsystem.OnChangedIterations

	form := widget.NewForm(
		widget.NewFormItem("Iterations", iterationsSlider),
	)

	return container.NewVBox(form)
}

func buildRightPane(lsystem *Lsystem) fyne.CanvasObject {
	lsystem.canvasImage.FillMode = canvas.ImageFillStretch

	return container.NewStack(lsystem.canvasImage)
}
