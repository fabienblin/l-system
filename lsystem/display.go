package lsystem

type BindableGrammarInterface interface {
	DisplayableInterface
	GrammarInterface
}

type DisplayableInterface interface {
	GetTurtle() Turtle
	SetTurtle(Turtle)
}

type Turtle struct {
	X, Y, Angle float64
}

type Displayable struct {
	Turtle     Turtle
}

func NewDisplayable() Displayable {
	return Displayable{
		Turtle: Turtle{},
	}
}

func (d *Displayable) GetTurtle() Turtle {
	return d.Turtle
}

func (d *Displayable) SetTurtle(turtle Turtle) {
	d.Turtle = turtle
}

