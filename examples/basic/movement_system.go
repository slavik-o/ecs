package main

import "github.com/slavik-o/ecs"

// MovementSystem updates positions based on velocities
type MovementSystem struct {
	requiredMask ecs.ComponentMask
}

// NewMovementSystem creates a new MovementSystem
func NewMovementSystem() *MovementSystem {
	// Require both position and velocity components
	return &MovementSystem{
		requiredMask: ecs.CreateComponentMask(COMPONENT_POSITION, COMPONENT_VELOCITY),
	}
}

// ComponentMask returns the component mask used to filter entities
func (s *MovementSystem) ComponentMask() ecs.ComponentMask {
	return s.requiredMask
}

// Update updates the positions of entities based on their velocities
func (s *MovementSystem) Update(dt float32, world *ecs.World) {
	// Get all entities with Position and Velocity components
	entities := world.GetEntitiesWithMask(s.requiredMask)

	for _, entity := range entities {
		// Get components
		pos := world.GetComponent(entity, COMPONENT_POSITION).(*Position)
		vel := world.GetComponent(entity, COMPONENT_VELOCITY).(*Velocity)

		// Update position based on velocity and delta time
		pos.X += vel.X * dt
		pos.Y += vel.Y * dt
	}
}
