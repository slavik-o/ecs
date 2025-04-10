package main

import (
	"fmt"

	"github.com/slavik-o/ecs"
)

// RenderSystem renders entities
type RenderSystem struct {
	requiredMask ecs.ComponentMask
}

// NewRenderSystem creates a new RenderSystem
func NewRenderSystem() *RenderSystem {
	return &RenderSystem{
		requiredMask: ecs.CreateComponentMask(
			COMPONENT_POSITION,
			COMPONENT_RENDERABLE,
		),
	}
}

// ComponentMask returns the component mask used to filter entities
func (s *RenderSystem) ComponentMask() ecs.ComponentMask {
	return s.requiredMask
}

// Update renders entities of the world
func (s *RenderSystem) Update(dt float32, world *ecs.World) {
	// Get all entities that have the required components
	entities := world.GetEntitiesWithMask(s.requiredMask)

	// Iterate over all entities that have the required components
	for _, entity := range entities {
		// Get the components of the entity
		position := world.GetComponent(entity, COMPONENT_POSITION).(*Position)
		renderable := world.GetComponent(entity, COMPONENT_RENDERABLE).(*Renderable)

		// Simulates rendering
		fmt.Printf("Rendering %s at %.0f,%.0f\n", renderable.Sprite, position.X, position.Y)
	}
}
