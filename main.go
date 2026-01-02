package main

import (
	"log"

	"main/grammar"
	"main/gui"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	grammar, errGrammar := grammar.NewAlgaeGrammar()
	if errGrammar != nil {
		log.Fatal(errGrammar.Error())
	}

	binder := gui.NewBinder(grammar)

	window := gui.BuildMainWindow(app, binder)
	window.ShowAndRun()
}
