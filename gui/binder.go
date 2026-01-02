package gui

import (
	"image"
	"image/color"
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
		canvasImage: canvas.NewImageFromImage(drawImage(grammar, grammar.GetAxiom())),
	}
}

func (b *Binder) OnChangedIterations(iterations float64) {
	b.grammar.SetTurtle(initTurtle())
	symbols := lsystem.IterateTranslationOnGrammar(int(iterations), b.grammar)
	b.canvasImage.Image = drawImage(b.grammar, symbols)
	b.canvasImage.Refresh()
}

func drawImage(grammar lsystem.BindableGrammarInterface, symbols string) image.Image {
	drawingContext := gg.NewContext(imageWidth, imageHeight)

	for _, symbol := range symbols {
		grammar.GetRule(symbol).DrawingFunc(drawingContext)
	}

	return drawingContext.Image()
}

func initTurtle() lsystem.Turtle {
	return lsystem.Turtle{
		X: float64(imageWidth)/2,
		Y: float64(imageHeight),
		Angle: .5,
	}
}