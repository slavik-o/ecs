package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Initialize Ebitengine
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("ecs")

	// Run the game
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
