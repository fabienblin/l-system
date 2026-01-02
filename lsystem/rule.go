package lsystem

import "github.com/fogleman/gg"

type DrawingFunction func(turtle *Turtle, ctx *gg.Context)

type RuleInterface interface {
	GetRune() rune
	GetTranslation() string
	// Drawing(turtle *Turtle, ctx *gg.Context)
}

type Rule struct {
	Symbol rune
	Translation string
	DrawingFunc DrawingFunction
}

func NewRule(symbol rune, translation string, drawingFunction DrawingFunction) Rule {
	return Rule{
		Symbol: symbol,
		Translation: translation,
		DrawingFunc: drawingFunction,
	}
}

func (r Rule) GetRune() rune {
	return r.Symbol
}

func (r Rule) GetTranslation() string {
	return r.Translation
}

// func (r Rule) Drawing(turtle *Turtle, ctx *gg.Context) {
// 	r.DrawingFunc(turtle, ctx)
// }

type Rules map[rune]Rule

func NewRules(rules ...Rule) Rules {
	var newRules Rules = make(Rules)
	for _, rule := range rules {
		newRules[rule.Symbol] = NewRule(rule.Symbol, rule.Translation, rule.DrawingFunc)
	}
	return newRules
}

func (rules Rules) GetRule(r rune) Rule {
	return rules[r]
}