package ecs

// World manages all entities, components and systems
type World struct {
	// Entity counter to generate unique entity IDs
	entityCounter Entity

	// Entities map to store all entities
	entities map[Entity]bool

	// Components slice to store all components
	components []ComponentID

	// Component stores map to store all component stores
	componentStores map[ComponentID]ComponentStore

	// Systems slice to store all systems
	systems []System

	// Entity masks map to store all entity masks
	entityMasks map[Entity]ComponentMask

	// Event manager to manage all events between systems
	EventManager *EventManager

	// State manager to manage entity states
	StateManager *StateManager
}

// NewWorld creates a new ECS world
func NewWorld() *World {
	return &World{
		entityCounter:   0,
		entities:        make(map[Entity]bool),
		components:      []ComponentID{},
		componentStores: make(map[ComponentID]ComponentStore),
		systems:         []System{},
		entityMasks:     make(map[Entity]ComponentMask),
		EventManager:    NewEventManager(),
		StateManager:    NewStateManager(),
	}
}

// RegisterComponentType registers a new component type and returns its ID
func (w *World) RegisterComponentType(id ComponentID) {
	w.componentStores[id] = &GenericComponentStore{
		components: make(map[Entity]interface{}),
	}
}

// NewEntity creates a new entity
func (w *World) NewEntity() Entity {
	w.entityCounter++

	w.entities[w.entityCounter] = true

	w.entityMasks[w.entityCounter] = 0 // No components initially

	return w.entityCounter
}

// RemoveEntity removes an entity and all its components
func (w *World) RemoveEntity(entity Entity) {
	if _, exists := w.entities[entity]; !exists {
		return
	}

	mask := w.entityMasks[entity]

	for _, id := range w.components {
		if (mask & (1 << id)) != 0 {
			w.componentStores[id].Remove(entity)
		}
	}

	w.StateManager.RemoveEntity(entity)
	delete(w.entities, entity)
	delete(w.entityMasks, entity)
}

// AddComponent adds a component to an entity
func (w *World) AddComponent(entity Entity, componentID ComponentID, component interface{}) {
	if store, exists := w.componentStores[componentID]; exists {
		store.Add(entity, component)

		w.entityMasks[entity] |= (1 << componentID)
	}
}

// RemoveComponent removes a component from an entity
func (w *World) RemoveComponent(entity Entity, componentID ComponentID) {
	if store, exists := w.componentStores[componentID]; exists {
		store.Remove(entity)

		w.entityMasks[entity] &= ^(1 << componentID)
	}
}

// GetComponent retrieves a component from an entity
func (w *World) GetComponent(entity Entity, componentID ComponentID) interface{} {
	if store, exists := w.componentStores[componentID]; exists {
		return store.Get(entity)
	}

	return nil
}

// HasComponent checks if an entity has a specific component
func (w *World) HasComponent(entity Entity, componentID ComponentID) bool {
	mask, exists := w.entityMasks[entity]

	if !exists {
		return false
	}

	return (mask & (1 << componentID)) != 0
}

// AddSystem adds a system to the world
func (w *World) AddSystem(system System) {
	w.systems = append(w.systems, system)
}

// Update updates all systems
func (w *World) Update(dt float32) error {
	if err := w.EventManager.Update(); err != nil {
		return err
	}

	if err := w.StateManager.Update(w, dt); err != nil {
		return err
	}

	for _, system := range w.systems {
		if err := system.Update(dt, w); err != nil {
			return err
		}
	}

	return nil
}

// GetEntitiesWithMask returns entities that match the component mask
func (w *World) GetEntitiesWithMask(mask ComponentMask) []Entity {
	var matchingEntities []Entity

	for entity := range w.entities {
		if (w.entityMasks[entity] & mask) == mask {
			matchingEntities = append(matchingEntities, entity)
		}
	}

	return matchingEntities
}
