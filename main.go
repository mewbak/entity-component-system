// Package main contains the entry point of the application.
package main

import (
	"github.com/davidsbond/game/pkg/game"
)

func main() {
	cnf, err := game.Load()

	if err != nil {
		panic(err)
	}

	gm := game.New(cnf)

	if err := gm.Start(); err != nil {
		panic(err)
	}
}
