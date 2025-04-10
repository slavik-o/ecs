package ecs

// GenericComponentStore is a concrete implementation of ComponentStore
type GenericComponentStore struct {
	// Components map to store all components
	components map[Entity]interface{}
}

// Initialize a new GenericComponentStore
func NewGenericComponentStore() *GenericComponentStore {
	return &GenericComponentStore{
		components: make(map[Entity]interface{}),
	}
}

// GenericComponentStore methods implementation
func (s *GenericComponentStore) Add(entity Entity, component interface{}) {
	s.components[entity] = component
}

// Remove removes a component from an entity
func (s *GenericComponentStore) Remove(entity Entity) {
	delete(s.components, entity)
}

// Get returns the component for an entity
func (s *GenericComponentStore) Get(entity Entity) interface{} {
	return s.components[entity]
}

// Has returns true if the entity has a component
func (s *GenericComponentStore) Has(entity Entity) bool {
	_, exists := s.components[entity]
	return exists
}
