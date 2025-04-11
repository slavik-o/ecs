package shared

import (
	"fmt"

	"github.com/slavik-o/ecs"
)

// CollisionSystem handles collision detection
type CollisionSystem struct {
	requiredMask ecs.ComponentMask
}

func NewCollisionSystem() *CollisionSystem {
	return &CollisionSystem{requiredMask: ecs.CreateComponentMask(COMPONENT_POSITION)}
}

func (s *CollisionSystem) ComponentMask() ecs.ComponentMask {
	return s.requiredMask
}

func (s *CollisionSystem) Update(dt float32, world *ecs.World) {
	entities := world.GetEntitiesWithMask(s.requiredMask)

	// Very simplistic collision detection for demonstration
	for i, entity1 := range entities {
		pos1 := world.GetComponent(entity1, COMPONENT_POSITION).(*Position)

		for j := i + 1; j < len(entities); j++ {
			entity2 := entities[j]
			pos2 := world.GetComponent(entity2, COMPONENT_POSITION).(*Position)

			// Check if entities are close enough to trigger collision
			dx := pos1.X - pos2.X
			dy := pos1.Y - pos2.Y
			distSquared := dx*dx + dy*dy

			if distSquared < 1.0 { // Arbitrary collision threshold
				// Publish collision event
				fmt.Printf("Collision detected between entities %d and %d\n", entity1, entity2)
				world.EventManager.Publish(&CollisionEvent{
					Entity1: entity1,
					Entity2: entity2,
				})
			}
		}
	}
}
