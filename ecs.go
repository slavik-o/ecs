package ecs

// Entity is just a unique identifier
type Entity uint32

// ComponentID represents the type of a component
type ComponentID uint64

// ComponentMask is a bitmask representing which components an entity has
type ComponentMask uint64

// System processes entities with specific components
type System interface {
	Update(dt float32, world *World)
	ComponentMask() ComponentMask
}

// ComponentStore is an interface for component storage
type ComponentStore interface {
	Add(entity Entity, component interface{})
	Remove(entity Entity)
	Get(entity Entity) interface{}
	Has(entity Entity) bool
}
