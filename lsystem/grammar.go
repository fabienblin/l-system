package lsystem

import (
	"log"
	"strings"
)

type GrammarInterface interface {
	GetAxiom() string
	GetRule(symbol rune) Rule
	GetRules() Rules
}

type Grammar struct {
	Axiom string
	Rules Rules
}

func NewGrammar(axiom string, rules Rules) (*Grammar, error) {
	grammar := &Grammar{
		Axiom: axiom,
		Rules: rules,
	}

	return grammar, nil
}

func (g *Grammar) GetAxiom() string {
	return g.Axiom
}

func (g *Grammar) GetRule(symbol rune) Rule {
	return g.Rules.GetRule(symbol)
}

func (g *Grammar) GetRules() Rules {
	return g.Rules
}

func IterateTranslationOnGrammar(i int, g GrammarInterface) string {
	translation := g.GetAxiom()
	for j := 0; j < i; j++ {
		log.Println(translation)
		translation = translateRule(translation, g)
	}

	return translation
}

func translateRule(previousTranslation string, g GrammarInterface) string {
	var translation strings.Builder
	for _, symbol := range previousTranslation {
		translation.WriteString(g.GetRule(symbol).GetTranslation())
	}

	return translation.String()
}
