package grammar

import (
	"main/lsystem"

	"github.com/fogleman/gg"
)

/**
 * AlgaeGrammar has 2 translatables A and B and no actionables
 * A translates to AB, and B translates to A
 * axiom is A
 */
type AlgaeGrammar struct {
	lsystem.Grammar
	lsystem.Displayable
}

func NewAlgaeGrammar() (*AlgaeGrammar, error) {
	algae := &AlgaeGrammar{}

	axiom := "A"
	ruleA := lsystem.NewRule('A', "AB", func(turtle lsystem.Turtle, ctx *gg.Context) {})

	ruleB := lsystem.NewRule('B', "A", func(turtle lsystem.Turtle, ctx *gg.Context) {})

	rules := lsystem.NewRules(ruleA, ruleB)

	grammar, errIntegrity := lsystem.NewGrammar(axiom, rules)
	if errIntegrity != nil {
		return nil, errIntegrity
	}

	algae.Grammar = *grammar

	return algae, nil
}
