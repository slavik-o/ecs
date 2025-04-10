package main

import (
	"fmt"
	"time"

	"github.com/slavik-o/ecs"
)

func main() {
	// Create a new world
	world := ecs.NewWorld()

	// Register systems
	world.AddSystem(NewRenderSystem())
	world.AddSystem(NewMovementSystem())

	// Register component types
	COMPONENT_POSITION = world.RegisterComponentType()
	COMPONENT_RENDERABLE = world.RegisterComponentType()
	COMPONENT_VELOCITY = world.RegisterComponentType()

	// Create player entity
	player := world.NewEntity()

	// Add components to player entity
	world.AddComponent(player, COMPONENT_RENDERABLE, &Renderable{Sprite: "player"})
	world.AddComponent(player, COMPONENT_POSITION, &Position{X: 10, Y: 10})
	world.AddComponent(player, COMPONENT_VELOCITY, &Velocity{X: 1, Y: 1})

	// Create enemy entity
	enemy := world.NewEntity()

	// Add components to enemy entity
	world.AddComponent(enemy, COMPONENT_RENDERABLE, &Renderable{Sprite: "enemy"})
	world.AddComponent(enemy, COMPONENT_POSITION, &Position{X: 20, Y: 20})
	world.AddComponent(enemy, COMPONENT_VELOCITY, &Velocity{X: -1, Y: -1})

	// Game loop simulation
	for i := range 10 {
		fmt.Printf("\n--- Frame %d ---\n", i+1)

		// Update with delta time of 1.0
		world.Update(1.0)

		time.Sleep(500 * time.Millisecond)
	}
}
