package shared

import "github.com/slavik-o/ecs"

// MovementSystem updates positions based on velocities
type MovementSystem struct {
	requiredMask ecs.ComponentMask
}

func NewMovementSystem() *MovementSystem {
	return &MovementSystem{
		requiredMask: ecs.CreateComponentMask(COMPONENT_POSITION, COMPONENT_VELOCITY),
	}
}

func (s *MovementSystem) ComponentMask() ecs.ComponentMask {
	return s.requiredMask
}

func (s *MovementSystem) Update(dt float32, world *ecs.World) {
	entities := world.GetEntitiesWithMask(s.requiredMask)

	for _, entity := range entities {
		pos := world.GetComponent(entity, COMPONENT_POSITION).(*Position)
		vel := world.GetComponent(entity, COMPONENT_VELOCITY).(*Velocity)

		// Update position based on velocity and delta time
		pos.X += vel.X * dt
		pos.Y += vel.Y * dt
	}
}
