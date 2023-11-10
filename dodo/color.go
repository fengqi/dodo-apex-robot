package dodo

import "math/rand"

var colors = []string{
	"grey",
	"red",
	"orange",
	"green",
	"blue",
	"purple",
	"default",
}

func RandColor() string {
	return colors[rand.Intn(len(colors))]
}
