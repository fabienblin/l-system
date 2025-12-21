package gui

import (
	"image"
	"image/color"
	"image/draw"
	"main/lsystem"

	"fyne.io/fyne/v2/canvas"
	"github.com/fogleman/gg"
)

var (
	backgroundColor = color.NRGBA{R: 200, G: 200, B: 200, A: 255}
)

const (
	imageWidth  int = 500
	imageHeight int = 500
)

/**
 * A Binder is a method that allows to bind a lsystem grammar to a visual representation
 */
type Binder struct {
	grammar     lsystem.BindableGrammarInterface
	canvasImage *canvas.Image
}

func NewBinder(grammar lsystem.BindableGrammarInterface) *Binder {
	return &Binder{
		grammar:     grammar,
		canvasImage: canvas.NewImageFromImage(drawImage(grammar)),
	}
}

func (b *Binder) OnChangedIterations(iterations float64) {
	lsystem.IterateTranslationOnGrammar(int(iterations), b.grammar)
	b.canvasImage.Image = drawImage(b.grammar)
	b.canvasImage.Refresh()
}

func drawImage(grammar lsystem.BindableGrammarInterface) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	draw.Draw(
		img,
		img.Bounds(),
		&image.Uniform{backgroundColor},
		image.Point{},
		draw.Src,
	)

	dc := gg.NewContextForRGBA(img)
	symbolImages := grammar.GetDisplayImages()
	for r := range grammar.GetTranslatables() {
		dc.DrawImageAnchored(symbolImages[r], imageWidth/2, imageHeight, 0.5, 0.5)
	}

	return img
}
