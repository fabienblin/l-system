package grammar

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
	lsystem.Grammar
	lsystem.Displayable
}

func NewAlgaeGrammar() (*AlgaeGrammar, error) {
	algae := &AlgaeGrammar{}

	initiator := "A"
	translatables := lsystem.Translatable{
		'A': func() string {
			return "AB"
		},
		'B': func() string {
			return "A"
		},
	}
	actionables := make(lsystem.Actionable)

	grammar, errIntegrity := lsystem.NewGrammar(initiator, translatables, actionables)
	if errIntegrity != nil {
		return nil, errIntegrity
	}

	displayables := lsystem.NewDisplayable(
		lsystem.SymbolImage{
			'A': algae.imageForA(),
			'B': algae.imageForB(),
		},
	)

	algae.Grammar = *grammar
	algae.Displayable = *displayables

	return algae, nil
}

func (g *AlgaeGrammar) imageForA() image.Image {
	dc := gg.NewContext(50, 50)
	dc.DrawCircle(25, 25, 25)
	dc.SetRGB(0, 0, 1)
	dc.Fill()
	return dc.Image()
}

func (g *AlgaeGrammar) imageForB() image.Image {
	dc := gg.NewContext(50, 50)
	dc.DrawCircle(25, 25, 25)
	dc.SetRGB(0, 1, 0)
	dc.Fill()
	return dc.Image()
}
