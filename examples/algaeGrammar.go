package examples

import (
	"image"
	"main/lsystem"

	"github.com/fogleman/gg"
)

/**
 * AlgaeGrammar has 2 translatables A and B and no actionables
 * A translates to AB, and B translates to A
 * initiator is A
 */
type AlgaeGrammar struct {
	initiator    string
	translatabes map[rune]lsystem.TranslatableRule
	actionables  map[rune]lsystem.ActionableRule
	X, Y float64
}

func NewAlgaeGrammar() (*AlgaeGrammar, error) {
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

	if errIntegrity := lsystem.VerifyIntegrity(grammar); errIntegrity != nil {
		return nil, errIntegrity
	}

	return grammar, nil
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

func (g *AlgaeGrammar) GetDisplayImages() map[rune]image.Image {
	displayImages := map[rune]image.Image{
		'A': imageForA(),
		'B': imageForB(),
	}

	return displayImages
}

func imageForA() image.Image {
	dc := gg.NewContext(50, 50)
	dc.DrawCircle(25, 25, 25)
	dc.SetRGB(0, 0, 1)
	dc.Fill()
	return dc.Image()
}

func imageForB() image.Image {
	dc := gg.NewContext(50, 50)
	dc.DrawCircle(25, 25, 25)
	dc.SetRGB(0, 1, 0)
	dc.Fill()
	return dc.Image()
}
