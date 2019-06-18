// Package main contains the entry point of the application.
package main

import (
	"github.com/davidsbond/game/internal/config"
	"github.com/davidsbond/game/internal/game"
)

func main() {
	cnf, err := config.Load()

	if err != nil {
		panic(err)
	}

	gm := game.New(cnf)

	if err := gm.Start(); err != nil {
		panic(err)
	}
}
