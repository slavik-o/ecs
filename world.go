package ecs

// World manages all entities, components and systems
type World struct {
	// Entity counter to generate unique entity IDs
	entityCounter Entity

	// Entities map to store all entities
	entities map[Entity]bool

	// Component stores map to store all component stores
	componentStores map[ComponentID]ComponentStore

	// Systems slice to store all systems
	systems []System

	// Entity masks map to store all entity masks
	entityMasks map[Entity]ComponentMask

	// Next component ID to generate unique component IDs
	nextComponentID ComponentID

	// Event manager to manage all events between systems
	EventManager *EventManager
}

// NewWorld creates a new ECS world
func NewWorld() *World {
	return &World{
		entityCounter:   0,
		entities:        make(map[Entity]bool),
		componentStores: make(map[ComponentID]ComponentStore),
		systems:         []System{},
		entityMasks:     make(map[Entity]ComponentMask),
		nextComponentID: 0,
		EventManager:    NewEventManager(),
	}
}

// RegisterComponentType registers a new component type and returns its ID
func (w *World) RegisterComponentType() ComponentID {
	// Get next component ID
	id := w.nextComponentID

	// Increment next component ID
	w.nextComponentID++

	// Initialize component store
	w.componentStores[id] = &GenericComponentStore{
		components: make(map[Entity]interface{}),
	}

	return id
}

// NewEntity creates a new entity
func (w *World) NewEntity() Entity {
	// Increment entity counter
	w.entityCounter++

	// Add entity to entities map
	w.entities[w.entityCounter] = true

	// Initialize entity mask
	w.entityMasks[w.entityCounter] = 0 // No components initially

	return w.entityCounter
}

// RemoveEntity removes an entity and all its components
func (w *World) RemoveEntity(entity Entity) {
	// Check if entity exists
	if _, exists := w.entities[entity]; !exists {
		return
	}

	// Remove entity from all component stores
	mask := w.entityMasks[entity]
	for id := ComponentID(0); id < w.nextComponentID; id++ {
		// Check if entity has component
		if (mask & (1 << id)) != 0 {
			// Remove entity from component store
			w.componentStores[id].Remove(entity)
		}
	}

	// Remove entity from entities map
	delete(w.entities, entity)

	// Remove entity from entity masks
	delete(w.entityMasks, entity)
}

// AddComponent adds a component to an entity
func (w *World) AddComponent(entity Entity, componentID ComponentID, component interface{}) {
	// Check if component store exists
	if store, exists := w.componentStores[componentID]; exists {
		// Add component to component store
		store.Add(entity, component)

		// Set the bit for this component
		w.entityMasks[entity] |= (1 << componentID)
	}
}

// RemoveComponent removes a component from an entity
func (w *World) RemoveComponent(entity Entity, componentID ComponentID) {
	// Check if component store exists
	if store, exists := w.componentStores[componentID]; exists {
		// Remove component from component store
		store.Remove(entity)

		// Clear the bit for this component
		w.entityMasks[entity] &= ^(1 << componentID)
	}
}

// GetComponent retrieves a component from an entity
func (w *World) GetComponent(entity Entity, componentID ComponentID) interface{} {
	// Check if component store exists
	if store, exists := w.componentStores[componentID]; exists {
		// Get component from component store
		return store.Get(entity)
	}

	// Return nil if component store does not exist
	return nil
}

// HasComponent checks if an entity has a specific component
func (w *World) HasComponent(entity Entity, componentID ComponentID) bool {
	// Check if entity mask exists
	mask, exists := w.entityMasks[entity]

	// Return false if entity mask does not exist
	if !exists {
		return false
	}

	// Return true if entity has component
	return (mask & (1 << componentID)) != 0
}

// AddSystem adds a system to the world
func (w *World) AddSystem(system System) {
	// Add system to systems
	w.systems = append(w.systems, system)
}

// Update updates all systems
func (w *World) Update(dt float32) {
	// Process events first
	w.EventManager.Update()

	// Update all systems
	for _, system := range w.systems {
		system.Update(dt, w)
	}
}

// GetEntitiesWithMask returns entities that match the component mask
func (w *World) GetEntitiesWithMask(mask ComponentMask) []Entity {
	var matchingEntities []Entity

	// Iterate over all entities
	for entity := range w.entities {
		// Check if entity's component mask has all the bits set in the required mask
		if (w.entityMasks[entity] & mask) == mask {
			// Add entity to matching entities
			matchingEntities = append(matchingEntities, entity)
		}
	}

	return matchingEntities
}
