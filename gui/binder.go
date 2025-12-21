package gui

import (
	"image"
	"image/color"
	"image/draw"

	"fyne.io/fyne/v2/canvas"
	"github.com/llgcode/draw2d/draw2dimg"
)

var (
	backgroundColor = color.NRGBA{R: 200, G: 200, B: 200, A: 255}
)

const (
	imageWidth  int = 500
	imageHeight int = 500
)

type Binder struct {
	canvasImage *canvas.Image
}

func NewBinder() *Binder {
	return &Binder{
		canvasImage: canvas.NewImageFromImage(drawImage()),
	}
}

func (b *Binder) OnChangedIterations(iterations float64) {
	b.canvasImage.Image = drawImage()
	b.canvasImage.Refresh()
}

func drawImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	draw.Draw(
		img,
		img.Bounds(),
		&image.Uniform{backgroundColor},
		image.Point{},
		draw.Src,
	)

	// stroke(img, color.Black, 5, 0,0,500,500)

	return img
}

func stroke(image *image.RGBA, c color.Color, width float64, x0 float64, y0 float64, x1 float64, y1 float64) {
	gc := draw2dimg.NewGraphicContext(image)
	gc.SetStrokeColor(c)
	gc.SetLineWidth(width)

	gc.MoveTo(x0, y0)
	gc.LineTo(x1, y1)
	gc.Stroke()
}
