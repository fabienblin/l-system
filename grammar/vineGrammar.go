package grammar

import (
	"image/color"
	"main/lsystem"
	"math"

	"github.com/fogleman/gg"
)

type VineGrammar struct {
	lsystem.DisplayableGrammar

	segment float64
	angle   float64
	stack   []lsystem.Turtle
}

func NewVineGrammar() (*VineGrammar, error) {
	g := &VineGrammar{
		segment: 8,
		angle:   25 * math.Pi / 180,
		stack:   []lsystem.Turtle{},
	}

	axiom := "F"

	ruleF := lsystem.NewRule('F', "FF-[-F+F+F]+[+F-F-F]", g.drawForward)
	rulePlus := lsystem.NewRule('+', "+", g.turnLeft)
	ruleMinus := lsystem.NewRule('-', "-", g.turnRight)
	rulePush := lsystem.NewRule('[', "[", g.pushState)
	rulePop := lsystem.NewRule(']', "]", g.popState)

	grammar, err := lsystem.NewDisplayableGrammar(axiom, ruleF, rulePlus, ruleMinus, rulePush, rulePop)
	if err != nil {
		return nil, err
	}

	g.DisplayableGrammar = *grammar

	return g, nil
}

func (g *VineGrammar) drawForward(ctx *gg.Context) {
	t := g.GetTurtle()

	nx := t.X + math.Cos(t.Angle)*g.segment
	ny := t.Y + math.Sin(t.Angle)*g.segment

	ctx.SetColor(color.White)
	ctx.DrawLine(t.X, t.Y, nx, ny)
	ctx.Stroke()

	t.X, t.Y = nx, ny
	g.SetTurtle(t)
}

func (g *VineGrammar) turnLeft(ctx *gg.Context) {
	t := g.GetTurtle()
	t.Angle += g.angle
	g.SetTurtle(t)
}

func (g *VineGrammar) turnRight(ctx *gg.Context) {
	t := g.GetTurtle()
	t.Angle -= g.angle
	g.SetTurtle(t)
}

func (g *VineGrammar) pushState(ctx *gg.Context) {
	t := g.GetTurtle()
	g.stack = append(g.stack, t)
}

func (g *VineGrammar) popState(ctx *gg.Context) {
	n := len(g.stack) - 1
	t := g.stack[n]
	g.stack = g.stack[:n]
	g.SetTurtle(t)
}
