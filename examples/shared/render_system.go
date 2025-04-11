package shared

import (
	"fmt"

	"github.com/slavik-o/ecs"
)

type RenderSystem struct {
	requiredMask ecs.ComponentMask
}

func NewRenderSystem() *RenderSystem {
	return &RenderSystem{
		requiredMask: ecs.CreateComponentMask(
			COMPONENT_POSITION,
			COMPONENT_RENDERABLE,
		),
	}
}

func (s *RenderSystem) ComponentMask() ecs.ComponentMask {
	return s.requiredMask
}

func (s *RenderSystem) Update(dt float32, world *ecs.World) {
	entities := world.GetEntitiesWithMask(s.requiredMask)

	for _, entity := range entities {
		position := world.GetComponent(entity, COMPONENT_POSITION).(*Position)
		renderable := world.GetComponent(entity, COMPONENT_RENDERABLE).(*Renderable)

		// Simulates rendering
		fmt.Printf("Rendering %s at %.0f,%.0f\n", renderable.Sprite, position.X, position.Y)
	}
}
