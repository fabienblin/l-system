package lsystem

import (
	"fmt"
	"image"
)

type TranslatableRule func() string
type ActionableRule func()

type Actionable map[rune]ActionableRule
type Translatable map[rune]TranslatableRule
type SymbolImage map[rune]image.Image

type GrammarInterface interface {
	GetInitiator() string
	GetTranslatables() Translatable
	GetActionables() Actionable
}

/**
 * Use this strcut by composition to easly implement GrammarInterface
 */
type Grammar struct {
	Initiator     string
	Translatables Translatable
	Actionables   Actionable
}

func NewGrammar(
	initiator string,
	translatables Translatable,
	actionables Actionable,
) (*Grammar, error) {
	grammar := &Grammar{
		Initiator:     initiator,
		Translatables: translatables,
		Actionables:   actionables,
	}

	if errIntegrity := VerifyIntegrity(grammar); errIntegrity != nil {
		return nil, errIntegrity
	}

	return grammar, nil
}

func (g *Grammar) GetInitiator() string {
	return g.Initiator
}

func (g *Grammar) GetTranslatables() Translatable {
	return g.Translatables
}

func (g *Grammar) GetActionables() Actionable {
	return g.Actionables
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
				g.GetActionables()[r]()
			} else if isTranslatableRune(r, g) {
				translated += g.GetTranslatables()[r]()
			}
		}
		previousTranslation = translated
	}

	return tree
}
