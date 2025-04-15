package main

import (
	"image/color"

	"examples/shared"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/slavik-o/ecs"
)

const (
	TILE_SIZE = 16
)

type Game struct {
	world *ecs.World
	// Mask to filter entities to render
	renderMask ecs.ComponentMask
}

func NewGame() *Game {
	world := ecs.NewWorld()

	// Register component types
	shared.COMPONENT_CONTROLLER = world.RegisterComponentType()
	shared.COMPONENT_POSITION = world.RegisterComponentType()
	shared.COMPONENT_RENDERABLE = world.RegisterComponentType()

	// Create player entity
	player := world.NewEntity()
	world.AddComponent(player, shared.COMPONENT_CONTROLLER, nil)
	world.AddComponent(player, shared.COMPONENT_RENDERABLE, &shared.Renderable{Sprite: "player"})
	world.AddComponent(player, shared.COMPONENT_POSITION, &shared.Position{X: 10, Y: 10})

	// Register systems
	world.AddSystem(NewControllerSystem(player))
	world.AddSystem(NewMovementSystem(world))

	// Initialize game
	return &Game{world: world, renderMask: ecs.CreateComponentMask(
		shared.COMPONENT_POSITION,
		shared.COMPONENT_RENDERABLE,
	)}
}

func (g *Game) Update() error {
	return g.world.Update(1.0 / 60.0) // Fixed delta time
}

func (g *Game) Draw(screen *ebiten.Image) {
	entities := g.world.GetEntitiesWithMask(g.renderMask)

	for _, entity := range entities {
		position := g.world.GetComponent(entity, shared.COMPONENT_POSITION).(*shared.Position)

		// Draw a rectangle at the entity's position
		vector.DrawFilledRect(screen,
			position.X*TILE_SIZE,
			position.Y*TILE_SIZE,
			TILE_SIZE,
			TILE_SIZE,
			color.RGBA{0xff, 0x00, 0x00, 0xff},
			false,
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
