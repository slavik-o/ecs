package shared

import "github.com/slavik-o/ecs"

// Component IDs for the components
const (
	COMPONENT_HEALTH ecs.ComponentID = iota
	COMPONENT_POSITION
	COMPONENT_RENDERABLE
	COMPONENT_VELOCITY
	COMPONENT_CONTROLLER
)

type Renderable struct {
	Sprite string
}

type Position struct {
	X, Y float32
}

type Velocity struct {
	X, Y float32
}

type Health struct {
	Current, Max int
}
