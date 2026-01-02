package lsystem

type DisplayableGrammar struct {
	Grammar
	Displayable
}

func NewDisplayableGrammar(axiom string, rules ...Rule) (*DisplayableGrammar, error) {
	g := &DisplayableGrammar{}
	
	newRules := NewRules(rules...)

	grammar, err := NewGrammar(axiom, newRules)
	if err != nil {
		return nil, err
	}

	g.Grammar = *grammar
	g.Displayable = *NewDisplayable()

	return g, nil
}
