package main

import (
	"examples/shared"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/slavik-o/ecs"
)

type ControllerSystem struct {
	controlledEntity ecs.Entity
	requiredMask     ecs.ComponentMask
}

func NewControllerSystem(controlledEntity ecs.Entity) *ControllerSystem {
	return &ControllerSystem{
		controlledEntity: controlledEntity,
		requiredMask: ecs.CreateComponentMask(
			shared.COMPONENT_CONTROLLER,
		),
	}
}

func (s *ControllerSystem) ComponentMask() ecs.ComponentMask {
	return s.requiredMask
}

func (s *ControllerSystem) Update(dt float32, world *ecs.World) error {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft):
		world.EventManager.Publish(&shared.MoveEvent{
			Entity:    s.controlledEntity,
			Direction: shared.DIRECTION_LEFT,
		})
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowRight):
		world.EventManager.Publish(&shared.MoveEvent{
			Entity:    s.controlledEntity,
			Direction: shared.DIRECTION_RIGHT,
		})
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowUp):
		world.EventManager.Publish(&shared.MoveEvent{
			Entity:    s.controlledEntity,
			Direction: shared.DIRECTION_UP,
		})
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowDown):
		world.EventManager.Publish(&shared.MoveEvent{
			Entity:    s.controlledEntity,
			Direction: shared.DIRECTION_DOWN,
		})
	}

	return nil
}
