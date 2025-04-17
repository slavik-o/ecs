package main

import (
	"fmt"
	"time"

	"examples/shared"

	"github.com/slavik-o/ecs"
)

func main() {
	// Create a new world
	world := ecs.NewWorld()

	// Register component types
	world.RegisterComponentType(shared.COMPONENT_POSITION)
	world.RegisterComponentType(shared.COMPONENT_RENDERABLE)
	world.RegisterComponentType(shared.COMPONENT_VELOCITY)

	// Create player entity
	player := world.NewEntity()

	// Add components to player entity
	world.AddComponent(player, shared.COMPONENT_RENDERABLE, &shared.Renderable{Sprite: "player"})
	world.AddComponent(player, shared.COMPONENT_POSITION, &shared.Position{X: 10, Y: 10})
	world.AddComponent(player, shared.COMPONENT_VELOCITY, &shared.Velocity{X: 1, Y: 1})

	// Create enemy entity
	enemy := world.NewEntity()

	// Add components to enemy entity
	world.AddComponent(enemy, shared.COMPONENT_RENDERABLE, &shared.Renderable{Sprite: "enemy"})
	world.AddComponent(enemy, shared.COMPONENT_POSITION, &shared.Position{X: 20, Y: 20})
	world.AddComponent(enemy, shared.COMPONENT_VELOCITY, &shared.Velocity{X: -1, Y: -1})

	// Register systems
	world.AddSystem(shared.NewRenderSystem())
	world.AddSystem(shared.NewMovementSystem())

	// Game loop simulation
	for i := range 10 {
		fmt.Printf("\n--- Frame %d ---\n", i+1)

		// Update with delta time of 1.0
		if err := world.Update(1.0); err != nil {
			panic(err)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
