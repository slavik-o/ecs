package ecs

// State represents a state that an entity can be in
type State interface {
	// Enter is called when entering this state
	Enter(entity Entity, world *World) error

	// Exit is called when exiting this state
	Exit(entity Entity, world *World) error

	// Update is called every frame while in this state
	Update(entity Entity, world *World, dt float32) error
}

// StateComponent represents the current state of an entity
type StateComponent struct {
	CurrentState  State
	PreviousState State
}

// StateManager manages entity states
type StateManager struct {
	// Map of entity to their state components
	states map[Entity]*StateComponent
}

// NewStateManager creates a new state manager
func NewStateManager() *StateManager {
	return &StateManager{
		states: make(map[Entity]*StateComponent),
	}
}

// SetState changes the state of an entity
func (sm *StateManager) SetState(entity Entity, world *World, newState State) error {
	stateComp, exists := sm.states[entity]
	if !exists {
		stateComp = &StateComponent{}
		sm.states[entity] = stateComp
	}

	// Call exit on current state if it exists
	if stateComp.CurrentState != nil {
		if err := stateComp.CurrentState.Exit(entity, world); err != nil {
			return err
		}
		stateComp.PreviousState = stateComp.CurrentState
	}

	// Set and enter new state
	stateComp.CurrentState = newState
	if err := newState.Enter(entity, world); err != nil {
		return err
	}

	return nil
}

// GetState returns the current state of an entity
func (sm *StateManager) GetState(entity Entity) State {
	if stateComp, exists := sm.states[entity]; exists {
		return stateComp.CurrentState
	}
	return nil
}

// GetPreviousState returns the previous state of an entity
func (sm *StateManager) GetPreviousState(entity Entity) State {
	if stateComp, exists := sm.states[entity]; exists {
		return stateComp.PreviousState
	}
	return nil
}

// Update updates all entity states
func (sm *StateManager) Update(world *World, dt float32) error {
	for entity, stateComp := range sm.states {
		if stateComp.CurrentState != nil {
			if err := stateComp.CurrentState.Update(entity, world, dt); err != nil {
				return err
			}
		}
	}
	return nil
}

// RemoveEntity removes an entity's state
func (sm *StateManager) RemoveEntity(entity Entity) {
	delete(sm.states, entity)
}
