package grammar

/**
 * AlgaeGrammar has 2 translatables A and B and no actionables
 * A translates to AB, and B translates to A
 * initiator is A
 */
type AlgaeGrammar struct {
	initiator    string
	translatabes map[rune]translatableRule
	actionables  map[rune]actionableRule
}

func NewAlgaeGrammar() *AlgaeGrammar {
	grammar := &AlgaeGrammar{
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

	if !VerifyIntegrity(grammar) {
		return nil
	}

	return grammar
}

func (g *AlgaeGrammar) getInitiator() string {
	return g.initiator
}

func (g *AlgaeGrammar) getTranslatables() map[rune]translatableRule {
	return g.translatabes
}

func (g *AlgaeGrammar) getActionables() map[rune]actionableRule {
	return g.actionables
}
