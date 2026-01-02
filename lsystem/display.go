package lsystem

type BindableGrammarInterface interface {
	DisplayableInterface
	GrammarInterface
}

type DisplayableInterface interface {
	GetTurtle() *Turtle
	SetIterations(int)
}

type Turtle struct {
	X, Y, Angle float64
}

/**
 * Use this strcut by composition to easly implement DisplayableInterface
 */
type Displayable struct {
	Turtle     *Turtle
	Iterations int
}

func NewDisplayable() *Displayable {
	return &Displayable{
		Turtle: &Turtle{},
	}
}

func (d *Displayable) GetTurtle() *Turtle {
	return d.Turtle
}

func (d *Displayable) SetIterations(iterations int) {
	d.Iterations = iterations
}
