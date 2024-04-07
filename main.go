package main

import (
	"fmt"
	"snake-game/game"
)

func main() {
	game := game.NewGame()
	go game.Before()

	for game.IsRunning() {
		game.Draw()
		game.Update()
	}

	fmt.Printf("\nYou scored %d points", game.Score())
}
