package main

import (
	"log"

	"main/gui"
	"main/lsystem"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	binder := gui.NewBinder()

	grammar := lsystem.NewAlgaeGrammar()
	iterations := lsystem.IterateTranslationOnGrammar(7, grammar)
	for i := 0; i < 7; i++ {
		log.Println(iterations[i])
	}

	window := gui.BuildMainWindow(app, binder)
	window.ShowAndRun()
}
