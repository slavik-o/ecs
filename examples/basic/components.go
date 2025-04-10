package main

import "github.com/slavik-o/ecs"

// Renderable is a component that can be rendered
type Renderable struct {
	Sprite string
}

// Position is a component that represents the position of an entity
type Position struct {
	X, Y float32
}

// Velocity is a component that represents the velocity of an entity
type Velocity struct {
	X, Y float32
}

// Component IDs for the components
var (
	COMPONENT_POSITION   ecs.ComponentID
	COMPONENT_RENDERABLE ecs.ComponentID
	COMPONENT_VELOCITY   ecs.ComponentID
)
