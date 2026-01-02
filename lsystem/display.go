package lsystem

const (
	TurtleMoveDistance float64 = 100
)

type DisplayableInterface interface {
	GetTurtle() Turtle
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
	Turtle     Turtle
	Iterations int
}

func NewDisplayable() *Displayable {
	return &Displayable{}
}

func (d *Displayable) GetTurtle() Turtle {
	return d.Turtle
}

func (d *Displayable) SetIterations(iterations int) {
	d.Iterations = iterations
}
