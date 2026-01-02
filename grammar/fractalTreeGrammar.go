package grammar

// import (
// 	"image"
// 	"main/lsystem"
// 	"math"

// 	"github.com/fogleman/gg"
// )

// /**
//  * translatables: 0, 1
//  * actionables: [, ]
//  * rules: (1 → 11), (0 → 1[0]0)
//  * axiom: "0"
//  */
// type FractalTreeGrammar struct {
// 	lsystem.Grammar
// 	lsystem.Displayable
// 	stackI int
// 	stack  []lsystem.Turtle
// }

// func NewFractalTreeGrammar() (*FractalTreeGrammar, error) {
// 	fractalTree := &FractalTreeGrammar{
// 		stack: make([]lsystem.Turtle, 0),
// 	}

// 	axiom := "0"
// 	translatables := lsystem.Translatable{
// 		'0': func() string { return "1[0]0" },
// 		'1': func() string { return "11" },
// 	}
// 	actionables := lsystem.Actionable{
// 		'[': fractalTree.funcForPush,
// 		']': fractalTree.funcForPop,
// 	}

// 	grammar, errIntegrity := lsystem.NewGrammar(axiom, translatables, actionables)
// 	if errIntegrity != nil {
// 		return nil, errIntegrity
// 	}

// 	displayables := lsystem.NewDisplayable(
// 		lsystem.SymbolImage{
// 			'0': fractalTree.imageFor0(),
// 			'1': fractalTree.imageFor1(),
// 		},
// 	)

// 	fractalTree.Grammar = *grammar
// 	fractalTree.Displayable = *displayables

// 	return fractalTree, nil
// }

// func (g *FractalTreeGrammar) imageFor0() image.Image {
// 	dc := gg.NewContext(50, 50)
// 	dc.DrawLine(25, 50, 25, 20)
// 	dc.DrawCircle(25, 10, 10)
// 	dc.SetRGB(0, 0, 1)
// 	dc.Fill()
// 	return dc.Image()
// }

// func (g *FractalTreeGrammar) imageFor1() image.Image {
// 	dc := gg.NewContext(50, 50)
// 	dc.DrawLine(25, 50, 25, 0)
// 	dc.SetRGB(0, 0, 1)
// 	dc.Fill()
// 	return dc.Image()
// }

// func (g *FractalTreeGrammar) funcForPop() {
// 	g.Turtle.Angle += math.Pi / 4

// }

// func (g *FractalTreeGrammar) funcForPush() {
// 	g.Turtle.Angle -= math.Pi / 4
// }
