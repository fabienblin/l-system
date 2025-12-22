package lsystem

const (
	TurtleMoveDistance float64 = 100
)

type DisplayableInterface interface {
	GetSymbolImages() SymbolImage
	GetTurtle() (Turtle)
	SetIterations(int)
}

type BindableGrammarInterface interface {
	DisplayableInterface
	GrammarInterface
}

type Turtle struct {
	X, Y, Angle float64
}

/**
 * Use this strcut by composition to easly implement DisplayableInterface
 */
type Displayable struct {
	Displayables SymbolImage
	Turtle       Turtle
	Iterations   int
}

func NewDisplayable(symbolImage SymbolImage) *Displayable {
	return &Displayable{
		Displayables: symbolImage,
	}
}

func (d *Displayable) GetSymbolImages() SymbolImage {
	return d.Displayables
}

func (d *Displayable) GetTurtlePosition() (Turtle) {
	return d.Turtle
}

func (d *Displayable) SetIterations(iterations int) {
	d.Iterations = iterations
}
