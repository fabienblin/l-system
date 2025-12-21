package main

import (
	"log"

	"main/grammar"
	"main/gui"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	binder := gui.NewBinder()

	algae := grammar.NewAlgaeGrammar()
	iterations := grammar.IterateTranslationOnGrammar(7, algae)
	for i := 0; i < 7; i++ {
		log.Println(iterations[i])
	}

	window := gui.BuildMainWindow(app, binder)
	window.ShowAndRun()
}
