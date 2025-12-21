package lsystem

import (
	"fmt"
)

type TranslatableRule func() string

type ActionableRule func() string

type GrammarInterface interface {
	GetInitiator() string
	GetTranslatables() map[rune]TranslatableRule
	GetActionables() map[rune]ActionableRule
}

func VerifyIntegrity(g GrammarInterface) error {
	for _, r := range g.GetInitiator() {
		actionable := isActionableRune(r, g)
		translatable := isTranslatableRune(r, g)
		if !actionable && !translatable {
			if !actionable {
				return fmt.Errorf("this grammar has an erroneous actionable rune %s", string(r))
			}
			if !translatable {
				return fmt.Errorf("this grammar has an erroneous translatable rune %s", string(r))
			}
		}
	}

	return nil
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
