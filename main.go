package main

import (
	"log"

	"main/grammar"
	"main/gui"
	"main/lsystem"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	grammar, errGrammar := grammar.NewAlgaeGrammar()
	if errGrammar != nil {
		log.Fatal(errGrammar.Error())
	}

	iterations := lsystem.IterateTranslationOnGrammar(7, grammar)
	for i := 0; i < 7; i++ {
		log.Println(iterations[i])
	}

	binder := gui.NewBinder(grammar)

	window := gui.BuildMainWindow(app, binder)
	window.ShowAndRun()
}
