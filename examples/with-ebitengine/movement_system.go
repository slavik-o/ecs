package main

import (
	"examples/shared"

	"github.com/slavik-o/ecs"
)

type MovementSystem struct {
	requiredMask ecs.ComponentMask
	world        *ecs.World
}

func NewMovementSystem(world *ecs.World) *MovementSystem {
	system := &MovementSystem{
		requiredMask: ecs.CreateComponentMask(
			shared.COMPONENT_CONTROLLER,
			shared.COMPONENT_POSITION,
			shared.COMPONENT_RENDERABLE,
		),
		world: world,
	}

	world.EventManager.Subscribe(shared.EVENT_MOVE, system.OnMove)

	return system
}

func (s *MovementSystem) ComponentMask() ecs.ComponentMask {
	return s.requiredMask
}

func (s *MovementSystem) Update(dt float32, world *ecs.World) {
	// do nothing
}

func (s *MovementSystem) OnMove(event ecs.Event) {
	moveEvent := event.(*shared.MoveEvent)

	entities := s.world.GetEntitiesWithMask(s.requiredMask)

	for _, entity := range entities {
		// Change position
		position := s.world.GetComponent(entity, shared.COMPONENT_POSITION).(*shared.Position)

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
	}
}
