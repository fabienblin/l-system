package grammar

type translatableRule func() string

type actionableRule func() string

type grammarInterface interface {
	getInitiator() string
	getTranslatables() map[rune]translatableRule
	getActionables() map[rune]actionableRule
}

func VerifyIntegrity(g grammarInterface) bool {
	for _, r := range g.getInitiator() {
		if !isActionableRune(r, g) && !isTranslatableRune(r, g) {
			return false
		}
	}

	return true
}

func isTranslatableRune(r rune, g grammarInterface) bool {
	_, translatable := g.getTranslatables()[r]
	if !translatable {
		return false
	}

	return true
}

func isActionableRune(r rune, g grammarInterface) bool {
	_, actionable := g.getActionables()[r]
	if !actionable {
		return false
	}

	return true
}

func IterateTranslationOnGrammar(i int, g grammarInterface) []string {
	tree := make([]string, i)

	previousTranslation := g.getInitiator()
	for j := 0; j < i; j++ {
		translated := ""
		tree[j] = previousTranslation
		for _, r := range previousTranslation {
			if isActionableRune(r, g) {
				translated += g.getActionables()[r]()
			} else if isTranslatableRune(r, g) {
				translated += g.getTranslatables()[r]()
			}
		}
		previousTranslation = translated
	}

	return tree
}
