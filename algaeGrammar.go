package main

/**
 * algaeGrammar has 2 translatables A and B and no actionables
 * A translates to AB, and B translates to A
 * initiator is A
 */
type algaeGrammar struct {
	initiator    string
	translatabes map[rune]translatableRule
	actionables  map[rune]actionableRule
}

func newAlgaeGrammar() *algaeGrammar {
	return &algaeGrammar{
		initiator: "A",
		translatabes: map[rune]translatableRule{
			'A': func() string {
				return "AB"
			},
			'B': func() string {
				return "A"
			},
		},
		actionables: make(map[rune]actionableRule),
	}
}

func (g *algaeGrammar) getInitiator() string {
	return g.initiator
}

func (g *algaeGrammar) getTranslatables() map[rune]translatableRule {
	return g.translatabes
}

func (g *algaeGrammar) getActionables() map[rune]actionableRule {
	return g.actionables
}
