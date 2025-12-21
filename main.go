package main

import (
	"log"

	"fyne.io/fyne/v2/app"
)

const (
	windowName   string  = "L System"
	windowWidth  float32 = 900
	windowHeight float32 = 600
)

func main() {
	app := app.New()
	lsystem := NewLsystem()
	
	grammar := newAlgaeGrammar()
	iterations := iterateTranslationOnGrammar(7, grammar)
	for i := 0; i < 7; i++ {
		log.Println(iterations[i])
	}

	window := buildMainWindow(app, lsystem)
	window.ShowAndRun()
}
