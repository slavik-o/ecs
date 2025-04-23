package main

import (
	"examples/shared"

	"github.com/slavik-o/ecs"
)

type MovementSystem struct{}

func NewMovementSystem(world *ecs.World) *MovementSystem {
	system := &MovementSystem{}

	world.EventManager.Subscribe(shared.EVENT_MOVE, system.onMove)

	return system
}

func (s *MovementSystem) ComponentMask() ecs.ComponentMask {
	return 0
}

func (s *MovementSystem) Update(dt float32, world *ecs.World) error {
	// do nothing
	return nil
}

func (s *MovementSystem) onMove(event ecs.Event, world *ecs.World) error {
	moveEvent := event.(*shared.MoveEvent)

	// Change position
	position := world.GetComponent(moveEvent.Entity, shared.COMPONENT_POSITION).(*shared.Position)

	switch moveEvent.Direction {
	case shared.DIRECTION_LEFT:
		position.X -= 1
	case shared.DIRECTION_RIGHT:
		position.X += 1
	case shared.DIRECTION_UP:
		position.Y -= 1
	case shared.DIRECTION_DOWN:
		position.Y += 1
	}

	return nil
}
