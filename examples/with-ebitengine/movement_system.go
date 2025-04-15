package main

import (
	"examples/shared"

	"github.com/slavik-o/ecs"
)

type MovementSystem struct {
	world *ecs.World
}

func NewMovementSystem(world *ecs.World) *MovementSystem {
	system := &MovementSystem{
		world: world,
	}

	world.EventManager.Subscribe(shared.EVENT_MOVE, system.OnMove)

	return system
}

func (s *MovementSystem) ComponentMask() ecs.ComponentMask {
	return 0
}

func (s *MovementSystem) Update(dt float32, world *ecs.World) error {
	// do nothing
	return nil
}

func (s *MovementSystem) OnMove(event ecs.Event) error {
	moveEvent := event.(*shared.MoveEvent)

	// Change position
	position := s.world.GetComponent(moveEvent.Entity, shared.COMPONENT_POSITION).(*shared.Position)

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
