package shared

import "github.com/slavik-o/ecs"

// Component IDs for the components
var (
	COMPONENT_HEALTH     ecs.ComponentID
	COMPONENT_POSITION   ecs.ComponentID
	COMPONENT_RENDERABLE ecs.ComponentID
	COMPONENT_VELOCITY   ecs.ComponentID
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
