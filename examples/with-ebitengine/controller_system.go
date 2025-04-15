package main

import (
	"examples/shared"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/slavik-o/ecs"
)

type ControllerSystem struct {
	requiredMask ecs.ComponentMask
}

func NewControllerSystem() *ControllerSystem {
	return &ControllerSystem{
		requiredMask: ecs.CreateComponentMask(
			shared.COMPONENT_CONTROLLER,
		),
	}
}

func (s *ControllerSystem) ComponentMask() ecs.ComponentMask {
	return s.requiredMask
}

func (s *ControllerSystem) Update(dt float32, world *ecs.World) {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft):
		world.EventManager.Publish(&shared.MoveEvent{
			Direction: shared.DIRECTION_LEFT,
		})
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowRight):
		world.EventManager.Publish(&shared.MoveEvent{
			Direction: shared.DIRECTION_RIGHT,
		})
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowUp):
		world.EventManager.Publish(&shared.MoveEvent{
			Direction: shared.DIRECTION_UP,
		})
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowDown):
		world.EventManager.Publish(&shared.MoveEvent{
			Direction: shared.DIRECTION_DOWN,
		})
	}
}
