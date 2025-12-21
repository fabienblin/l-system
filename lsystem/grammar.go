package lsystem

type TranslatableRule func() string

type ActionableRule func() string

type GrammarInterface interface {
	GetInitiator() string
	GetTranslatables() map[rune]TranslatableRule
	GetActionables() map[rune]ActionableRule
}

func VerifyIntegrity(g GrammarInterface) bool {
	for _, r := range g.GetInitiator() {
		if !isActionableRune(r, g) && !isTranslatableRune(r, g) {
			return false
		}
	}

	return true
}

func isTranslatableRune(r rune, g GrammarInterface) bool {
	_, translatable := g.GetTranslatables()[r]
	if !translatable {
		return false
	}

	return true
}

func isActionableRune(r rune, g GrammarInterface) bool {
	_, actionable := g.GetActionables()[r]
	if !actionable {
		return false
	}

	return true
}

func IterateTranslationOnGrammar(i int, g GrammarInterface) []string {
	tree := make([]string, i)

	previousTranslation := g.GetInitiator()
	for j := 0; j < i; j++ {
		translated := ""
		tree[j] = previousTranslation
		for _, r := range previousTranslation {
			if isActionableRune(r, g) {
				translated += g.GetActionables()[r]()
			} else if isTranslatableRune(r, g) {
				translated += g.GetTranslatables()[r]()
			}
		}
		previousTranslation = translated
	}

	return tree
}
