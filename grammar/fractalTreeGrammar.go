package grammar

import (
	"image/color"
	"main/lsystem"
	"math"

	"github.com/fogleman/gg"
)

/**
 * symbols: 0, 1, [, ]
 * rules: (1 → 11), (0 → 1[0]0), ([ -> left & push), (] -> right & pop)
 * axiom: "0"
 */
type FractalTreeGrammar struct {
	lsystem.DisplayableGrammar

	segmentSize   float64
	rotationAngle float64
	stack         []lsystem.Turtle
}

func NewFractalTreeGrammar() (*FractalTreeGrammar, error) {
	fractalTree := &FractalTreeGrammar{}
	fractalTree.segmentSize = 5
	fractalTree.rotationAngle = math.Pi / 4

	axiom := "0"

	rule0 := lsystem.NewRule('0', "1[0]0", fractalTree.Rule0DrawingFunction)

	rule1 := lsystem.NewRule('1', "11", fractalTree.Rule1DrawingFunction)

	ruleLeft := lsystem.NewRule('[', "[", fractalTree.RuleLeftDrawingFunction)

	ruleRight := lsystem.NewRule(']', "]", fractalTree.RuleRightDrawingFunction)

	grammar, errIntegrity := lsystem.NewDisplayableGrammar(axiom, rule0, rule1, ruleLeft, ruleRight)
	if errIntegrity != nil {
		return nil, errIntegrity
	}

	fractalTree.DisplayableGrammar = grammar

	return fractalTree, nil
}

func (g *FractalTreeGrammar) forward(ctx *gg.Context) {
	t := g.GetTurtle()

	nx := t.X + math.Cos(t.Angle)*g.segmentSize
	ny := t.Y + math.Sin(t.Angle)*g.segmentSize

	ctx.SetColor(color.White)
	ctx.DrawLine(t.X, t.Y, nx, ny)
	ctx.Stroke()

	t.X, t.Y = nx, ny
	g.SetTurtle(t)
}

func (g *FractalTreeGrammar) Rule0DrawingFunction(ctx *gg.Context) {
	g.forward(ctx)
}

func (g *FractalTreeGrammar) Rule1DrawingFunction(ctx *gg.Context) {
	g.forward(ctx)
}

func (g *FractalTreeGrammar) RuleLeftDrawingFunction(ctx *gg.Context) {
	t := g.GetTurtle()
	g.stack = append(g.stack, t)
	t.Angle += g.rotationAngle
	g.SetTurtle(t)
}

func (g *FractalTreeGrammar) RuleRightDrawingFunction(ctx *gg.Context) {
	n := len(g.stack) - 1
	t := g.stack[n]
	g.stack = g.stack[:n]

	t.Angle -= g.rotationAngle
	g.SetTurtle(t)
}
