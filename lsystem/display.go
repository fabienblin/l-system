package lsystem

import "image"

type DisplayableInterface interface {
	GetDisplayImages() map[rune]image.Image
}

type BindableGrammarInterface interface {
	DisplayableInterface
	GrammarInterface
}