package shared

import "github.com/slavik-o/ecs"

// Define event types
const (
	EVENT_COLLISION ecs.EventType = iota
	EVENT_ENTITY_DIED
	EVENT_HEALTH_CHANGED
)

// CollisionEvent is triggered when two entities collide
type CollisionEvent struct {
	Entity1, Entity2 ecs.Entity
}

func (e CollisionEvent) Type() ecs.EventType {
	return EVENT_COLLISION
}

// EntityDiedEvent is triggered when an entity's health reaches zero
type EntityDiedEvent struct {
	Entity ecs.Entity
}

func (e EntityDiedEvent) Type() ecs.EventType {
	return EVENT_ENTITY_DIED
}

// HealthChangedEvent is triggered when an entity's health changes
type HealthChangedEvent struct {
	Entity         ecs.Entity
	PreviousHealth int
	NewHealth      int
}

func (e HealthChangedEvent) Type() ecs.EventType {
	return EVENT_HEALTH_CHANGED
}
