package dodo

import "math/rand"

var themes = []string{
	"grey",
	"red",
	"orange",
	"yellow",
	"green",
	"indigo",
	"blue",
	"purple",
	"black",
	"default",
}

func RandTheme() string {
	return themes[rand.Intn(len(themes))]
}
