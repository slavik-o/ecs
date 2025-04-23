package main

import (
	"fmt"

	"examples/shared"

	"github.com/slavik-o/ecs"
)

// IdleState represents when an entity is standing still
type IdleState struct{}

func (s *IdleState) Enter(entity ecs.Entity, world *ecs.World) error {
	renderable := world.GetComponent(entity, shared.COMPONENT_RENDERABLE).(*shared.Renderable)

	fmt.Printf("%s entered Idle state\n", renderable.Sprite)
	return nil
}

func (s *IdleState) Exit(entity ecs.Entity, world *ecs.World) error {
	renderable := world.GetComponent(entity, shared.COMPONENT_RENDERABLE).(*shared.Renderable)

	fmt.Printf("%s exited Idle state\n", renderable.Sprite)
	return nil
}

func (s *IdleState) Update(entity ecs.Entity, world *ecs.World, dt float32) error {
	// In idle state, entity doesn't move
	velocity := world.GetComponent(entity, shared.COMPONENT_VELOCITY).(*shared.Velocity)
	velocity.X = 0
	velocity.Y = 0
	return nil
}

// MovingState represents when an entity is moving
type MovingState struct{}

func (s *MovingState) Enter(entity ecs.Entity, world *ecs.World) error {
	renderable := world.GetComponent(entity, shared.COMPONENT_RENDERABLE).(*shared.Renderable)

	fmt.Printf("%s entered Moving state\n", renderable.Sprite)
	return nil
}

func (s *MovingState) Exit(entity ecs.Entity, world *ecs.World) error {
	renderable := world.GetComponent(entity, shared.COMPONENT_RENDERABLE).(*shared.Renderable)

	fmt.Printf("%s exited Moving state\n", renderable.Sprite)
	return nil
}

func (s *MovingState) Update(entity ecs.Entity, world *ecs.World, dt float32) error {
	// In moving state, entity moves in its current direction
	velocity := world.GetComponent(entity, shared.COMPONENT_VELOCITY).(*shared.Velocity)
	position := world.GetComponent(entity, shared.COMPONENT_POSITION).(*shared.Position)

	// Update position based on velocity
	position.X += velocity.X * dt
	position.Y += velocity.Y * dt

	// Bounce off walls (simple boundary checking)
	if position.X <= 0 || position.X >= 100 {
		velocity.X = -velocity.X
	}
	if position.Y <= 0 || position.Y >= 100 {
		velocity.Y = -velocity.Y
	}

	return nil
}
