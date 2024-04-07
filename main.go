package main

import (
	"fmt"
	"snake-game/game"
	"time"
)

func main() {
	game := game.NewGame()
	go game.Before()

	for game.IsRunning() {
		game.Draw()
		time.Sleep(150 * time.Millisecond)
		game.Update()
	}

	fmt.Printf("\nYou scored %d points", game.Score())
}
