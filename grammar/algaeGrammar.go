package grammar

import (
	"image/color"
	"main/lsystem"

	"github.com/fogleman/gg"
)

/**
 * AlgaeGrammar has 2 translatables A and B and no actionables
 * A translates to AB, and B translates to A
 * axiom is A
 */
type AlgaeGrammar struct {
	lsystem.DisplayableGrammar
}

func NewAlgaeGrammar() (*AlgaeGrammar, error) {
	const (
		circleRadius float64 = -5
		moveY        float64 = circleRadius * 2
	)

	algae := &AlgaeGrammar{}

	axiom := "A"
	
	ruleA := lsystem.NewRule('A', "AB", func(ctx *gg.Context) {
		turtle := algae.GetTurtle()
		turtle.Y += moveY
		algae.SetTurtle(turtle)
		ctx.DrawCircle(turtle.X, turtle.Y, circleRadius)
		ctx.SetColor(color.White)
		ctx.Fill()
	})

	ruleB := lsystem.NewRule('B', "A", func(ctx *gg.Context) {
		turtle := algae.GetTurtle()
		turtle.Y += moveY
		algae.SetTurtle(turtle)
		ctx.DrawCircle(turtle.X, turtle.Y, circleRadius)
		ctx.SetColor(color.White)
		ctx.Stroke()
	})

	grammar, errIntegrity := lsystem.NewDisplayableGrammar(axiom, ruleA, ruleB)
	if errIntegrity != nil {
		return nil, errIntegrity
	}

	algae.DisplayableGrammar = *grammar

	return algae, nil
}
