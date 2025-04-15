package shared

import (
	"fmt"

	"github.com/slavik-o/ecs"
)

// HealthSystem handles health-related logic
type HealthSystem struct {
	requiredMask ecs.ComponentMask
	// World is needed to subscribe to events
	world *ecs.World
}

func NewHealthSystem(world *ecs.World) *HealthSystem {
	system := &HealthSystem{
		requiredMask: ecs.CreateComponentMask(COMPONENT_HEALTH),
		world:        world,
	}

	// Subscribe to events
	world.EventManager.Subscribe(EVENT_HEALTH_CHANGED, system.handleHealthChanged)
	world.EventManager.Subscribe(EVENT_ENTITY_DIED, system.handleEntityDied)

	return system
}

func (s *HealthSystem) ComponentMask() ecs.ComponentMask {
	return s.requiredMask
}

func (s *HealthSystem) Update(dt float32, world *ecs.World) error {
	// Could implement health regeneration or other ongoing processes here
	return nil
}

func (s *HealthSystem) handleHealthChanged(event ecs.Event) error {
	healthEvent := event.(*HealthChangedEvent)
	fmt.Printf("Entity %d health changed from %d to %d\n",
		healthEvent.Entity, healthEvent.PreviousHealth, healthEvent.NewHealth)

	return nil
}

func (s *HealthSystem) handleEntityDied(event ecs.Event) error {
	deathEvent := event.(*EntityDiedEvent)
	fmt.Printf("Entity %d has died!\n", deathEvent.Entity)

	return nil
}
