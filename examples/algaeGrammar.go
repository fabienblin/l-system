package examples

import "main/lsystem"

/**
 * AlgaeGrammar has 2 translatables A and B and no actionables
 * A translates to AB, and B translates to A
 * initiator is A
 */
type AlgaeGrammar struct {
	initiator    string
	translatabes map[rune]lsystem.TranslatableRule
	actionables  map[rune]lsystem.ActionableRule
}

func NewAlgaeGrammar() *AlgaeGrammar {
	grammar := &AlgaeGrammar{
		initiator: "A",
		translatabes: map[rune]lsystem.TranslatableRule{
			'A': func() string {
				return "AB"
			},
			'B': func() string {
				return "A"
			},
		},
		actionables: make(map[rune]lsystem.ActionableRule),
	}

	if !lsystem.VerifyIntegrity(grammar) {
		return nil
	}

	return grammar
}

func (g *AlgaeGrammar) GetInitiator() string {
	return g.initiator
}

func (g *AlgaeGrammar) GetTranslatables() map[rune]lsystem.TranslatableRule {
	return g.translatabes
}

func (g *AlgaeGrammar) GetActionables() map[rune]lsystem.ActionableRule {
	return g.actionables
}
